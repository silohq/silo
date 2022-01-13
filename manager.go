package silo

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

type manager struct {
	conn    *bolt.Tx
	buckets map[string]*bolt.Bucket
}

func createManager(conn *bolt.DB) (*manager, error) {
	tx, err := conn.Begin(true)
	if err != nil {
		return nil, fmt.Errorf("failed to init transaction")
	}

	return &manager{
		conn:    tx,
		buckets: make(map[string]*bolt.Bucket),
	}, nil
}

func (m *manager) createbkt(name string) error {
	_, ok := m.buckets[name]
	if !ok {
		bkt, err := m.conn.CreateBucket([]byte(name))
		if err != nil {
			return fmt.Errorf("failed to create bucket")
		}
		m.buckets[name] = bkt
	}
	return nil
}

func (m *manager) createchildbkt(graph string) error {
	nodes := strings.Split(graph, ".")
	if len(nodes) < 1 {
		return fmt.Errorf("invalid parent key")
	}
	root := nodes[0]
	if !m.exists(root) {
		m.createbkt(root)
	}

	for i := 0; i < len(nodes); i++ {
		key := strings.Join(nodes[0:i+1], ".")
		base := strings.Join(nodes[0:i], ".")
		bkt, ok := m.parent(base)
		if ok {
			tmp, err := bkt.CreateBucketIfNotExists([]byte(nodes[i]))
			if err != nil {
				return err
			}
			m.save(key, tmp)
		}
	}

	return nil
}

func (m *manager) parent(name string) (*bolt.Bucket, bool) {
	if val, ok := m.buckets[name]; ok {
		return val, true
	}
	return nil, false
}

func (m *manager) exists(name string) bool {
	if _, ok := m.buckets[name]; !ok {
		return false
	}

	return true
}

func (m *manager) all() map[string]*bolt.Bucket {
	return m.buckets
}

func (m *manager) save(key string, bkt *bolt.Bucket) {
	m.buckets[key] = bkt
}

// TODO: fix silent fail
func (m *manager) insert(tree map[string]interface{}) {
	bkts := m.all()
	id := uuid.New().String()
	for k, v := range tree {
		if bkt, ok := bkts[k]; ok {
			typ := fmt.Sprintf("%T", v)
			if typ == "string" {
				bkt.Put([]byte(id), []byte(v.(string)))
			}
		}
	}
}

func (m *manager) find(key string, match interface{}, tree map[string]interface{}) {
	bkt, ok := m.parent(key)
	if !ok {
		return
	}

	csr := bkt.Cursor()
	for id, v := csr.First(); id != nil; id, v = csr.Next() {
		if string(v) == match {
			for k := range tree {
				bkt, ok := m.parent(k)
				if ok {
					tree[k] = bkt.Get(id)
				}
			}
			return
		}
	}
}

func (m *manager) delete(key string, match interface{}, tree map[string]interface{}) {
	bkt, ok := m.parent(key)
	if !ok {
		return
	}

	csr := bkt.Cursor()
	for id, v := csr.First(); id != nil; id, v = csr.Next() {
		if string(v) == match {
			for k := range tree {
				bkt, ok := m.parent(k)
				if ok {
					bkt.Delete(id)
				}
			}
			return
		}
	}
}

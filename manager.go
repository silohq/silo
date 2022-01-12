package silo

import (
	"fmt"
	"log"
	"strings"

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
			log.Printf("exisyts %v", key)
			tmp, err := bkt.CreateBucketIfNotExists([]byte(nodes[i]))
			if err != nil {
				log.Printf("failed to create child bucket %s", err)
				return err
			}
			m.save(key, tmp)
		}
	}

	return nil
}

func (m *manager) parent(name string) (*bolt.Bucket, bool) {
	log.Printf("name %s", name)
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
	log.Printf("key %s", key)
	m.buckets[key] = bkt
}

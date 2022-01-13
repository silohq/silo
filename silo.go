package silo

import (
	"errors"
	"fmt"
	"strings"

	bolt "go.etcd.io/bbolt"
)

type Silo struct {
	conn   *bolt.DB
	config *Config
	root   *Tree
	mgr    *manager
}

type Config struct {
	DocPath string
	DBPath  string
}

func new(conf *Config) (*Silo, error) {
	db, err := bolt.Open(conf.DBPath, 0666, nil)
	if err != nil {
		return nil, errors.New("cannot start: failed to init db")
	}

	tree, err := ParseFromFile(conf.DocPath)
	if err != nil {
		return nil, fmt.Errorf("doc definition failed: %s", err)
	}

	mgr, err := createManager(db)
	if err != nil {
		return nil, fmt.Errorf("manager create failed: %s", err)
	}

	structread(mgr, tree)

	return &Silo{
		config: conf,
		conn:   db,
		root:   tree,
		mgr:    mgr,
	}, nil
}

func New(conf *Config) (*Silo, error) {
	return new(conf)
}

// create parses the command template and inserts into database
func (s *Silo) Create(command string, payload map[string]interface{}) {
	history := make(map[string]interface{})
	nestedmapaccess(payload, history)
	s.mgr.insert(history)
}

func (s *Silo) Find(command string, match interface{}, dest map[string]interface{}) error {
	history := make(map[string]interface{})
	nestedmapaccess(dest, history)
	s.mgr.find(command, match, history)
	for k, v := range history {
		keys := strings.Split(k, ".")
		if len(keys) > 1 {
			_, err := mapreverse(dest, v, keys)
			if err != nil {
				return fmt.Errorf("lookup failed %s", err)
			}
		}
	}
	return nil
}

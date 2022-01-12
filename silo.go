package silo

import (
	"errors"
	"fmt"
	"log"

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

	StructRead(mgr, tree)
	log.Printf("all %v", mgr.all())

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
func Create(command string, payload map[string]interface{}) {

}

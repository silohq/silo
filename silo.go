package silo

import (
	"errors"
	"log"

	bolt "go.etcd.io/bbolt"
)

type Silo struct {
	conn   *bolt.DB
	config *Config
	root   *Tree
}

type Config struct {
	DocPath string
	DBPath  string
}

func new(conf *Config) (*Silo, error) {
	db, err := bolt.Open(conf.DBPath, 0666, nil)
	if err != nil {
		log.Printf("cannot start db: %s", err)
		return nil, errors.New("cannot start: failed to init db")
	}

	tree, err := ParseFromFile(conf.DocPath)
	if err != nil {
		log.Printf("failed to parse doc definition: %s", err)
	}

	return &Silo{
		config: conf,
		conn:   db,
		root:   tree,
	}, nil
}

func New(conf *Config) (*Silo, error) {
	return new(conf)
}

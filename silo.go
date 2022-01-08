package silo

import (
	"errors"
	"log"

	bolt "go.etcd.io/bbolt"
)

type Silo struct {
	conn *bolt.DB
}

func New(path string) (*Silo, error) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		log.Printf("cannot start %s", err)
		return nil, errors.New("cannot start: failed to init db")
	}

	return &Silo{
		conn: db,
	}, nil
}

func (s *Silo) Parse(filename string) (*Tree, error) {
	return ParseFromFile(filename)
}

// func (s *Silo) Manager() *Manager {
// 	return worker.New(s.conn)
// }

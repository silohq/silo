package silo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func Parse(in []byte) error {
	dec := createDecoder(in)

	for dec.More() {
		node := make(map[string]interface{})
		err := dec.Decode(&node)
		if err != nil {
			log.Printf("something went wrong %s", err)
			return errors.New("failed to decode node")
		}

	}

	return nil
}

func createDecoder(in []byte) *json.Decoder {
	reader := bytes.NewReader(in)
	return json.NewDecoder(reader)
}

func ParseFromFile(path string) (*Tree, error) {
	dec := createFileDecoder(path)
	node := Tree{}

	for dec.More() {
		err := dec.Decode(&node)
		if err != nil {
			log.Printf("something went wrong %s", err)
			return nil, fmt.Errorf("failed to parse from file")
		}
	}

	return &node, nil
}

func createFileDecoder(path string) *json.Decoder {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file %s", err)
	}

	return json.NewDecoder(file)
}

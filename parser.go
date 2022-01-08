package silo

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
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

		log.Printf("Item %v", node)
	}

	return nil
}

func createDecoder(in []byte) *json.Decoder {
	reader := bytes.NewReader(in)
	return json.NewDecoder(reader)
}

package silo

import (
	"encoding/json"
	"log"
	"os"
)

func Parse(path string) {
	dec := createDecoder(path)

	for dec.More() {
		node := make(map[string]interface{})
		err := dec.Decode(&node)
		if err != nil {
			log.Printf("something went wrong %s", err)
		}

		log.Printf("Item %v", node)
	}
}

func createDecoder(path string) *json.Decoder {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file %s", err)
	}

	return json.NewDecoder(file)
}

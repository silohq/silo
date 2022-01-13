package silo

import (
	"fmt"
	"log"
)

func nestedmapaccess(in map[string]interface{}, tree map[string][]string) {
	for k, v := range in {
		typ := fmt.Sprintf("%T", v)
		tree[k] = []string{""}
		if typ == "map[string]interface {}" {
			deepaccess(k, v.(map[string]interface{}), tree)
		}
	}

	log.Printf("all %v", tree)
}

func deepaccess(parent string, in map[string]interface{}, tree map[string][]string) {
	for k, v := range in {
		typ := fmt.Sprintf("%T", v)
		tree[parent] = append(tree[parent], k)
		if typ == "map[string]interface {}" {
			deepaccess(parent, v.(map[string]interface{}), tree)
		}
	}
}

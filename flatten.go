package silo

import (
	"fmt"
)

func flatten(in map[string]interface{}, tree map[string]interface{}) {
	for k, v := range in {
		typ := fmt.Sprintf("%T", v)
		tree[k] = []string{""}
		if typ == "map[string]interface {}" {
			deepaccess(k, v.(map[string]interface{}), tree)
		}
	}
}

func deepaccess(parent string, in map[string]interface{}, tree map[string]interface{}) {
	for k, v := range in {
		typ := fmt.Sprintf("%T", v)
		if typ == "map[string]interface {}" {
			key := fmt.Sprintf("%s.%s", parent, k)
			deepaccess(key, v.(map[string]interface{}), tree)
		}

		if typ == "string" {
			key := fmt.Sprintf("%s.%s", parent, k)
			tree[key] = v
		}
	}
}

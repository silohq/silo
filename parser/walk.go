package parser

import (
	"log"
	"reflect"
)

func ReadMap(in map[string]interface{}) {
	for _, v := range in {
		// log.Printf("k hhh %v, val: %v", k, v)
		if ismap(v) {
			ReadMap(v.(map[string]interface{}))
		}
	}
}

func ismap(in interface{}) bool {
	tmp := reflect.ValueOf(in)
	log.Printf("here %v", tmp.Kind())
	return tmp.Kind() == reflect.Map
}

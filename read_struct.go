package silo

import (
	"reflect"
	"strings"
)

func StructRead(mgr *manager, in interface{}) {
	val := extract(in)
	parent := val.FieldByName("Parent").String()
	typ := val.FieldByName("Type").String()

	if parent != "" && parent == "root" {
		mgr.createbkt(val.FieldByName("Label").String())
	}

	if parent != "" {
		if typ == "nested-node" {
			mgr.createchildbkt(parent)
		}

		if typ == "flat-node" {
			label := val.FieldByName("Label").String()
			keys := strings.Split(parent, ".")
			keys = append(keys, label)
			merged := strings.Join(keys[:], ".")
			mgr.createchildbkt(merged)
		}
	}

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		//currently only checking for slice type
		if f.Kind() == reflect.Slice {
			for j := 0; j < f.Len(); j++ {
				StructRead(mgr, f.Index(j).Interface())
			}
		}
	}
}

func extract(in interface{}) reflect.Value {
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

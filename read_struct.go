package silo

import (
	"reflect"
)

func StructRead(in interface{}) {
	val := extract(in)
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		//currently only checking for slice type
		if f.Kind() == reflect.Slice {
			for j := 0; j < f.Len(); j++ {
				StructRead(f.Index(j).Interface())
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

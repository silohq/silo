package silo

import (
	"reflect"
)

func ReadStruct(st interface{}) {
	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			ReadStruct(f.Interface())
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				ReadStruct(f.Index(j).Interface())
			}
		case reflect.String:
			// fmt.Printf("%v=%v\n", val.Type().Field(i).Name, val.Field(i).Interface())
		}
	}
}

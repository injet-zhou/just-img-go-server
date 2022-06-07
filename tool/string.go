package tool

import (
	"reflect"
	"strings"
)

func TrimFields(s interface{}) interface{} {
	if s == nil {
		return nil
	}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.NumField() == 0 {
		return s
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.String {
			f.SetString(strings.Trim(f.String(), " "))
		}
	}
	return s
}

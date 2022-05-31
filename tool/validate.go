package tool

import "reflect"

func IsStructEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			IsStructEmpty(f.Interface())
		} else if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				return true
			}
			f = f.Elem()
			if f.Kind() == reflect.Struct {
				IsStructEmpty(f.Interface())
			}
		} else if f.Kind() == reflect.Slice {
			if f.IsNil() {
				return true
			}
			if f.Len() == 0 {
				return true
			}
			for j := 0; j < f.Len(); j++ {
				e := f.Index(j)
				if e.Kind() == reflect.Struct {
					IsStructEmpty(e.Interface())
				}
			}
		} else if f.Kind() == reflect.Map {
			if f.IsNil() {
				return true
			}
			if f.Len() == 0 {
				return true
			}
			for _, k := range f.MapKeys() {
				e := f.MapIndex(k)
				if e.Kind() == reflect.Struct {
					IsStructEmpty(e.Interface())
				}
			}
		} else if f.Kind() == reflect.String {
			if f.Len() == 0 {
				return true
			}
		} else if f.Kind() == reflect.Int {
			if f.Int() == 0 {
				return true
			}
		}
	}
	return false
}

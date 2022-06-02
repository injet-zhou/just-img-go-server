package tool

import (
	"reflect"
)

// IsStructEmpty 判断struct是否为空
func IsStructEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.NumField() == 0 {
		return true
	}
	isEmpty := false
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			isEmpty = IsStructEmpty(f.Interface())
			if isEmpty {
				return isEmpty
			}
		} else if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				return true
			}
			f = f.Elem()
			if f.Kind() == reflect.Struct || f.Kind() == reflect.Slice || f.Kind() == reflect.Map {
				isEmpty = IsStructEmpty(f.Interface())
				if isEmpty {
					return isEmpty
				}
			} else if f.Len() == 0 || f.Int() == 0 || f.Float() == 0 {
				return true
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
				if e.Kind() == reflect.Struct || e.Kind() == reflect.Slice || e.Kind() == reflect.Map {
					isEmpty = IsStructEmpty(e.Interface())
					if isEmpty {
						return isEmpty
					}
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
				if e.Kind() == reflect.Struct || e.Kind() == reflect.Slice || e.Kind() == reflect.Map {
					isEmpty = IsStructEmpty(e.Interface())
					if isEmpty {
						return isEmpty
					}
				} else if e.Len() == 0 || e.Int() == 0 || e.Float() == 0 {
					return true
				}
			}
		} else if f.Kind() == reflect.String {
			if f.Len() == 0 {
				return true
			}
		} else if f.Kind() == reflect.Int || f.Kind() == reflect.Int64 || f.Kind() == reflect.Int32 || f.Kind() == reflect.Int16 || f.Kind() == reflect.Int8 {
			if f.Int() == 0 {
				return true
			}
		} else if f.Kind() == reflect.Uint || f.Kind() == reflect.Uint64 || f.Kind() == reflect.Uint32 || f.Kind() == reflect.Uint16 || f.Kind() == reflect.Uint8 {
			if f.Uint() == 0 {
				return true
			}
		} else if f.Kind() == reflect.Float32 || f.Kind() == reflect.Float64 {
			if f.Float() == 0 {
				return true
			}
		}
	}
	return false
}

package tpls

import "reflect"

func init() {
	funcs["isSet"] = func(name string, data interface{}) bool {
		v := reflect.ValueOf(data)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Map {
			return v.MapIndex(reflect.ValueOf(name)).IsValid()
		}
		if v.Kind() != reflect.Struct {
			return false
		}
		return v.FieldByName(name).IsValid()
	}
}

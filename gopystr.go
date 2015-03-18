package gopystr

import (
	"fmt"
	"reflect"
	"strconv"
)

func str(obj interface{}) (s string, err error) {

	switch v := reflect.ValueOf(obj); v.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		s = strconv.FormatFloat(v.Float(), 'f', 12, 64)
	case reflect.Bool:
		if v.Bool() {
			s = "True"
		} else {
			s = "False"
		}
	case reflect.String:
		s = v.String()
	case reflect.Ptr:
		if ss, err := str(v.Pointer()); err != nil {
			s = ss
		}
	default:
		err = fmt.Errorf("Unknow type")
	}
	return
}

// Str method of Python implement by Golang
// Copyright (C) 2015 Meng Zhuo <mengzhuo1203@gmail.com>
// License can be found in the LICENSE file

package gopystr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Str(obj interface{}) (s string) {

	v := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		v = v.Elem()
	}

	switch typ.Kind() {

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
		v := v.Elem()
		s = Str(v)
	case reflect.Map:
		if typ.Key().Kind() != reflect.String {
			s = "{}"
		} else {
			var buf bytes.Buffer
			buf.WriteByte('{')

			for i, k := range v.MapKeys() {
				if i > 0 {
					buf.Write([]byte(", "))
				}
				buf.WriteString(quote(k.String(), '\''))
				buf.WriteByte(':')
				switch v.MapIndex(k).Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16,
					reflect.Int32, reflect.Int64, reflect.Uint,
					reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					buf.WriteString(Str(v.MapIndex(k).Interface()))
				default:
					buf.WriteString(quote(Str(v.MapIndex(k).Interface()), '\''))
				}
			}
			buf.WriteByte('}')
			s = buf.String()
		}
	case reflect.Slice:
		var buf bytes.Buffer
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.Write([]byte(", "))
			}
			switch v.Index(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64, reflect.Uint,
				reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				buf.WriteString(Str(v.Index(i).Interface()))
			default:
				buf.WriteString(quote(Str(v.Index(i).Interface()), '\''))
			}
		}
		buf.WriteByte(']')
		s = buf.String()

	case reflect.Struct:

		var buf bytes.Buffer
		buf.WriteByte('{')

		for i := 0; i < typ.NumField(); i++ {
			if i > 0 {
				buf.Write([]byte(", "))
			}
			buf.WriteString(quote(typ.Field(i).Name, '\''))
			buf.WriteByte(':')
			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64, reflect.Uint,
				reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				buf.WriteString(Str(v.Field(i).Interface()))
			default:
				buf.WriteString(quote(Str(v.Field(i).Interface()), '\''))
			}
		}
		buf.WriteByte('}')
		s = buf.String()
	default:
		s = fmt.Sprintf("%qv", obj)
	}
	return
}

func quote(str string, tag rune) string {
	return string(tag) + str + string(tag)
}

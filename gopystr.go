// Str method of Python implement by Golang
// Copyright (C) 2015 Meng Zhuo <mengzhuo1203@gmail.com>
// License can be found in the LICENSE file

package gopystr

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// Make Golang instance into string formatted as Python
//
// Examples:
//
//		gopystr.Str("Hi") // Hi
//		gopystr.Str(1)    // 1
//
//		// Map with key by string also supported
//		m := *&map[string]int{"A":1, "B":2}
//		gopystr.Str(m) // {'A': 1, 'B': 2}
//
//		// Struct also supported
//		type Oyster struct {
//			Closed bool
//      }
//      o := &Oyster{true}
//      gopystr.Str(o) // {"Closed":True}
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
			keys := v.MapKeys()
			sort.Sort(ByKey(keys))

			for i, k := range keys {
				if i > 0 {
					buf.Write([]byte(", "))
				}
				buf.WriteString(quote(k.String(), '\''))
				buf.Write([]byte(": "))
				switch v.MapIndex(k).Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16,
					reflect.Int32, reflect.Int64, reflect.Uint,
					reflect.Uint8, reflect.Uint16, reflect.Uint32,
					reflect.Uint64, reflect.Bool, reflect.Map, reflect.Struct:
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
				reflect.Uint8, reflect.Uint16, reflect.Uint32,
				reflect.Uint64, reflect.Bool, reflect.Map, reflect.Struct:
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
			buf.Write([]byte(": "))
			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64, reflect.Uint,
				reflect.Uint8, reflect.Uint16, reflect.Uint32,
				reflect.Uint64, reflect.Bool, reflect.Map, reflect.Struct:
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

type ByKey []reflect.Value

func (b ByKey) Len() int {
	return len(b)
}
func (b ByKey) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByKey) Less(i, j int) bool {
	return b[i].String() < b[j].String()
}

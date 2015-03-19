package gopystr

import (
	"testing"
)

func TestInt(t *testing.T) {
	int_slice := *&[]interface{}{int(1), int8(1), int16(1), int32(1), int64(1)}
	for _, num := range int_slice {
		if r := Str(num); r != "1" {
			t.Fatal("Int conv failed with:", r, num)
		}
	}
	minus_int_slice := *&[]interface{}{int(-1), int8(-1), int16(-1), int32(-1), int64(-1)}
	for _, num := range minus_int_slice {
		if r := Str(num); r != "-1" {
			t.Fatal("Int conv failed with:", r, num)
		}
	}
}

func TestIntPtr(t *testing.T) {
	a := 1
	if r := Str(&a); r != "1" {
		t.Fatal("IntPtr conv failed with:", r, a)
	}
	b := "TEST"
	if r := Str(&b); r != "TEST" {
		t.Fatal("StringPtr conv failed with:", r, a)
	}
}
func TestUInt(t *testing.T) {
	int_slice := *&[]interface{}{uint(1), uint8(1), uint16(1), uint32(1), uint64(1), 1}
	for _, num := range int_slice {
		if r := Str(num); r != "1" {
			t.Fatal("UInt conv failed with:", r, num)
		}
	}
}

/* FIXME BUGGY FLOAT
func TestFloat(t *testing.T) {
	float_slice := *&[]interface{}{float32(-0.230), float64(-0.230), -0.230, -0.23}
	for i, num := range float_slice {
		if r := Str(num); r[:5] != "-0.23" {
			t.Fatal("Float conv failed with:", i, r, num)
		}
	}
}
*/

func TestBool(t *testing.T) {
	if v := Str(true); v != "True" {
		t.Fatal("Bool conv failed with true")
	}
	if v := Str(false); v != "False" {
		t.Fatal("Bool conv failed with False")
	}
}

func TestMapStringKeyIntValue(t *testing.T) {
	d := *&map[string]int{"A": 1, "B": 2}
	if v := Str(d); v != "{\"A\":1, \"B\":2}" {
		t.Error("Map conv failed with False", d, v)
	}
}

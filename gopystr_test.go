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
	if v := Str(d); v != "{'A':1, 'B':2}" {
		t.Error("Map conv failed with False", d, v)
	}
}
func TestMapStringKeyStringValue(t *testing.T) {
	d := *&map[string]string{"A": "NOT A", "B": "NOT B"}
	if v := Str(d); v != "{'A':'NOT A', 'B':'NOT B'}" {
		t.Error("Map conv failed with False", d, v)
	}
}
func TestMapStringKeyStringValuePtr(t *testing.T) {
	d := &map[string]string{"A": "NOT A", "B": "NOT B"}
	if v := Str(d); v != "{'A':'NOT A', 'B':'NOT B'}" {
		t.Error("Map conv failed with False", d, v)
	}
}

func TestSliceStringValue(t *testing.T) {

	d := *&[]string{"A", "B"}

	if v := Str(d); v != "['A', 'B']" {
		t.Error("conv failed", d, v)
	}
}

func TestStruct(t *testing.T) {
	type TT struct {
		S string
		I int
		M map[string]int
	}

	var d TT
	d.S = "A"
	d.I = 1
	d.M = *&map[string]int{"M1": 1, "M2": 2}

	if v := Str(d); v != "{'S':'A', 'I':1, 'M':'{'M1':1, 'M2':2}'}" {
		t.Error("conv failed", d, v)
	}

}

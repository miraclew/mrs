package util

import (
	"fmt"
	r "reflect"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Format("2006-01-02 15:04:05"))
}

func TestMakeRandomString(t *testing.T) {
	// v1 := makeRandomString(10)
	// v2 := makeRandomString(10)
	// if v1 == v2 {
	// 	t.Fail()
	// }

	// v3 := makeRandomString(40)
	// fmt.Println(v3)
	// if len(v3) != 40 {
	// 	t.Fail()
	// }
}

type S struct {
	I   int
	Str string
}

type error string

func (e error) String() string { return string(e) }

func main() {
	s := S{5, "Bunnies"}
	m, _ := StructToMap(s)

	s1 := S{}
	MapToStruct(m, &s1)

	fmt.Printf("%+v\n%v\n%+v\n", s, m, s1)
}

func StructToMap(val interface{}) (mapVal map[string]interface{}, ok bool) {
	// indirect so function works with both structs and pointers to them
	structVal, ok := r.Indirect(r.NewValue(val)).(*r.StructValue)
	if !ok {
		return
	}

	typ := structVal.Type().(*r.StructType)
	mapVal = make(map[string]interface{})

	for i := 0; i < typ.NumField(); i++ {
		field := structVal.Field(i)
		if field.CanSet() {
			mapVal[typ.Field(i).Name] = field.Interface()
		}
	}

	return
}

func MapToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
	structVal, ok := r.Indirect(r.NewValue(val)).(*r.StructValue)
	if !ok {
		return
	}

	for name, elem := range mapVal {
		structVal.FieldByName(name).SetValue(r.NewValue(elem))
	}

	return
}

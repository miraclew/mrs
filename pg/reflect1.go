package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func inspect(f interface{}) map[string]string {

	m := make(map[string]string)
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		f := valueField.Interface()
		val := reflect.ValueOf(f)
		m[typeField.Name] = val.String()
	}

	return m
}

func dump(m map[string]string) {

	for k, v := range m {
		fmt.Printf("%s : %s\n", k, v)
	}
}

func main1() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	a := inspect(f)

	dump(a)
}

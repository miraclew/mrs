package main

import (
	"fmt"
	"reflect"
)

type Req1 struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

type PushHandler interface {
	NewChannel(subsId []int64) (channelId int64, err error)
	PushToUser(userId int64, message interface{}) (err error)
	PushToChannel(chanelId int64, message interface{}) (err error)
}

func inspect(f interface{}) map[string]string {

	m := make(map[string]string)
	val := reflect.TypeOf(f).Elem()

	for i := 0; i < val.NumMethod(); i++ {
		method := val.Method(i)
		fmt.Printf("%#v\n\n", method)
		funcs := method.Type
		fmt.Printf("%s\n", funcs)
		for j := 0; j < funcs.NumIn(); j++ {
			fmt.Printf("In: %s \n", funcs.In(j).String())
		}

		for k := 0; k < funcs.NumOut(); k++ {
			fmt.Printf("Out: %s \n", funcs.Out(k).String())
		}
	}

	return m
}

func dump(m map[string]string) {

	for k, v := range m {
		fmt.Printf("%s : %s\n", k, v)
	}
}

func main() {
	var f *PushHandler

	inspect(f)

	// dump(a)
}

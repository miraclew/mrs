package main

import (
	// "fmt"
	"os"
	"text/template"
)

type MProto struct {
	Types      string
	Interfaces uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	//tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	tmpl, err := template.ParseFiles("t.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

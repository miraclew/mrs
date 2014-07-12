package api

import (
	// "fmt"
	"github.com/miraclew/restful"
)

type Hello struct {
	Message string
}

type UserController struct {
	restful.ApiController
}

func (this *UserController) Get() {
	this.Data = &Hello{Message: "hello"}
}

func (this *UserController) Post() {
	username := this.Request.PostFormValue("username")
	// password := this.Request.PostFormValue("password")

	this.Response.Data = struct {
		Id      int
		UseName string
	}{
		123,
		username,
	}
}

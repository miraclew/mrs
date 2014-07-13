package api

import (
	"fmt"
	"github.com/miraclew/mrs/missle"
	"github.com/miraclew/restful"
)

type UserController struct {
	restful.ApiController
}

func (this *UserController) Post() {
	values := this.Request.PostForm
	fmt.Printf("Post: %#v \n", values)

	action := values.Get("a")

	if len(action) == 0 {
		this.Data = response(-1, "action is nil")
		return
	}

	if action == "register" {
		username := values.Get("username")
		password := values.Get("password")

		user := &missle.User{UserName: username, Password: password}
		user.Save()
		this.Data = response(0, nil)
		return
	} else if action == "autoRegister" {

	}
}

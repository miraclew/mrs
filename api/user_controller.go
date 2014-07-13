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

		if len(username) < 2 { //
			this.Data = response(ERR_INVALID_ARGS, "username too short")
			return
		}

		if len(password) < 3 { //
			this.Data = response(ERR_INVALID_ARGS, "password too short")
			return
		}

		// check exists
		uid := missle.GetUidByUserName(username)
		if uid != 0 {
			this.Data = response(ERR_INVALID_ARGS, "username is taken")
			return
		}

		user := &missle.User{UserName: username, Password: password}
		user.Save()
		this.Data = response(0, RegisterResponse{user.Uid})
		return
	} else if action == "autoRegister" {

	} else {
		this.Data = response(ERR_UNKNOWN_ACTION, action)
		return
	}
}

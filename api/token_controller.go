package api

import (
	"fmt"
	"github.com/miraclew/mrs/missle"
	"github.com/miraclew/restful"
)

type TokenController struct {
	restful.ApiController
}

func (this *TokenController) Post() {
	values := this.Request.PostForm
	fmt.Printf("TokenController Post: %#v \n", values)

	// action := values.Get("a")
	username := values.Get("username")
	password := values.Get("password")

	user := missle.FindUserByCredential(username, password)
	if user == nil {
		this.Data = response(ERR_INVALID_CREDENTIAL, "")
		return
	}

	token, err := missle.MakeToken(user.Uid)
	if err != nil {
		this.Data = response(ERR_INTERAL_ERROR, "make token failed")
		return
	}

	this.Data = &TokenResponse{user.Uid, token}
}

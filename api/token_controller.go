package api

import (
	"github.com/miraclew/restful"
)

type TokenController struct {
	restful.ApiController
}

func (this *TokenController) Get() {
	this.Response.Data = "TokenController"
}

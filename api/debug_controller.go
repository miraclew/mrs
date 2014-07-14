package api

import (
	"github.com/miraclew/restful"
)

type DebugController struct {
	restful.ApiController
}

func (this *DebugController) Get() {
	this.Response.Data = "DebugController"
}

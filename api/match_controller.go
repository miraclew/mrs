package api

import (
	// "fmt"
	"github.com/miraclew/restful"
)

type MatchController struct {
	restful.ApiController
}

func (this *MatchController) Get() {

}

func (this *MatchController) Post() {
	username := this.Request.PostFormValue("username")
	// password := this.Request.PostFormValue("password")
}

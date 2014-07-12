package api

import (
	// "fmt"
	"github.com/miraclew/mrs/missle"
	"github.com/miraclew/restful"
	"strconv"
)

type MatchController struct {
	restful.ApiController
}

func (this *MatchController) Post() {
	action := this.Request.PostFormValue("a")
	playerId := 1

	if action == "enter" {
		enter(playerId)
	}

	matchId := strconv.ParseInt(this.Request.PostFormValue("matchId"), 10, 64)
	match := missle.GetMatch(matchId)

	if action == "playerMove" {
		x := strconv.ParseFloat(this.Request.PostFormValue("x"), 32)
		y := strconv.ParseFloat(this.Request.PostFormValue("y"), 32)
		err := match.PlayerMove(playerId, missle.Point{x, y})

		this.Data = response(0, nil)
	}

	if action == "playerFire" {
		x := strconv.ParseFloat(this.Request.PostFormValue("x"), 32)
		y := strconv.ParseFloat(this.Request.PostFormValue("y"), 32)
		velocity := missle.Point{x, y}
		pos := missle.Point{0, 0}
		match.PlayerFire(playerId, pos, velocity)
		this.Data = response(0, nil)
	}

	if action == "targetHit" {
		p1 := strconv.ParseInt(this.Request.PostFormValue("p1"), 10, 64)
		p2 := strconv.ParseInt(this.Request.PostFormValue("p2"), 10, 64)
		health := strconv.ParseInt(this.Request.PostFormValue("p2"), 10, 32)

		match.PlayerHealth(p2, -health) // minus p2
		this.Data = response(0, nil)
	}
}

func targetHit(p1 int64, p2 int64, health int) {

}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func response(code int, data interface{}) *Response {
	r := &Response{
		code,
		data,
	}

	return r
}

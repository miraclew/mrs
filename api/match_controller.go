package api

import (
	"fmt"
	"github.com/miraclew/mrs/missle"
	"github.com/miraclew/restful"
	"strconv"
)

type MatchController struct {
	restful.ApiController
}

func (this *MatchController) Post() {
	values := this.Request.PostForm
	fmt.Printf("Post: %#v \n", values)

	action := values.Get("a")
	token := values.Get("token")

	if (len(action) == 0) || (len(token) == 0) {
		this.Data = response(-1, "action or token is nil")
		return
	}

	playerId, err := missle.GetUidByToken(token)
	if err != nil {
		this.Data = response(-1, err.Error())
		return
	}

	if action == "enter" {
		game := missle.GetGame()
		var pusher missle.Pusher
		pusher = &missle.PusherMock{}
		game.Pusher = pusher

		err := game.PlayerEnter(playerId)
		if err != nil {
			this.Data = response(-1, err.Error())
			return
		}

		this.Data = response(0, nil)
		return
	}

	matchId, _ := strconv.ParseInt(values.Get("matchId"), 0, 64)
	match := missle.GetMatch(matchId)
	if match == nil {
		this.Data = response(-1, "matchId is nil")
		return
	}

	if action == "playerMove" {
		x, _ := strconv.ParseFloat(values.Get("x"), 32)
		y, _ := strconv.ParseFloat(values.Get("y"), 32)
		err := match.PlayerMove(playerId, missle.Point{float32(x), float32(y)})

		if err != nil {
			this.Data = response(0, err.Error())
		} else {
			this.Data = response(0, nil)
		}
		return
	}

	if action == "playerFire" {
		x, _ := strconv.ParseFloat(values.Get("x"), 32)
		y, _ := strconv.ParseFloat(values.Get("y"), 32)
		velocity := missle.Point{float32(x), float32(y)}
		pos := missle.Point{0, 0}
		match.PlayerFire(playerId, pos, velocity)
		this.Data = response(0, nil)
		return
	}

	if action == "playerAttack" {
		p1, _ := strconv.ParseInt(values.Get("p1"), 0, 64)
		p2 := playerId
		damage, _ := strconv.Atoi(values.Get("damage"))

		match.PlayerAttack(p1, p2, damage)
		this.Data = response(0, nil)
		return
	}
}

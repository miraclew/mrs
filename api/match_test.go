package api

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestA(t *testing.T) {
	t.Skip("...")
	playerEnter("aa")
	playerEnter("bb")

	playerMove("aa", "0.2", "0.3")
	playerFire("aa", "0.2", "0.3")
	playerAttack("bb", "1", "30")
}

func TestUser(t *testing.T) {
	v := url.Values{}
	v.Set("a", "register")
	v.Set("username", "miraclew2")
	v.Set("password", "123")

	uri, _ := url.Parse("localhost:8080/match")
	req := &http.Request{PostForm: v, URL: uri}

	c := &UserController{}
	c.Init(nil, req)
	c.Post()

	fmt.Printf("response: %#v", c.Data)
}

func TestToken(t *testing.T) {
	t.Skip("...")
	v := url.Values{}
	v.Set("username", "miraclew")
	v.Set("password", "123")

	uri, _ := url.Parse("localhost:8080/match")
	req := &http.Request{PostForm: v, URL: uri}

	c := &TokenController{}
	c.Init(nil, req)
	c.Post()

	fmt.Printf("response: %#v", c.Data)
}

func sendPost(v url.Values) {
	uri, _ := url.Parse("localhost:8080/match")
	req := &http.Request{PostForm: v, URL: uri}

	c := &MatchController{}
	c.Init(nil, req)
	c.Post()

	fmt.Printf("response: %#v", c.Data)
}

func playerEnter(token string) {
	v := url.Values{}
	v.Set("token", token)
	v.Set("a", "enter")

	sendPost(v)
}

func playerMove(token string, x string, y string) {
	v := url.Values{}
	v.Set("token", token)
	v.Set("a", "playerMove")
	v.Set("matchId", "1") // this should got in ws
	v.Set("x", x)
	v.Set("y", y)

	sendPost(v)
}

func playerFire(token string, x string, y string) {
	v := url.Values{}
	v.Set("token", token)
	v.Set("a", "playerFire")
	v.Set("matchId", "1") // this should got in ws
	v.Set("x", x)
	v.Set("y", y)

	sendPost(v)
}

// token is p2's
func playerAttack(token string, p1 string, damage string) {
	v := url.Values{}
	v.Set("token", token)
	v.Set("a", "playerAttack")
	v.Set("matchId", "1") // this should got in ws
	v.Set("p1", p1)
	v.Set("damage", damage)

	sendPost(v)
}

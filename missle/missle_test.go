package missle

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	t.Skip("...")
	// var x float32
	// x = 1.0 * 2
	// y := x * 2
	// fmt.Printf("%f, %f", x, y)
	points := MakeKeyPoints(16)
	for _, p := range points {
		fmt.Printf("%f, %f\n", p.X, p.Y)
	}
}

func TestB(t *testing.T) {
	t.Skip("...")
	var pusher Pusher
	pusher = &PusherMock{}
	match := NewMatch([]int64{1, 2}, pusher)
	match.Begin()

	if match.State != STATE_PLAYING {
		t.Error("STATE error")
	}

	match.PlayerMove(1, Point{0.44, 0.34})
	match.PlayerFire(2, Point{0.5, 0.6}, Point{0.1, 0.2})
	match.PlayerHealth(1, -10)
	match.PlayerHealth(2, +10)

	match.PlayerHealth(1, -90)
}

func TestC(t *testing.T) {
	var pusher Pusher
	pusher = &PusherMock{}

	game := GetGame()
	game.Pusher = pusher
	game.PlayerEnter(1)
	game.PlayerEnter(2)
	game.PlayerEnter(3)
	game.PlayerEnter(4)
}

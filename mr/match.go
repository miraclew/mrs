package mr

import (
	"fmt"
	"io"
	"log"
)

const (
	StateCreated = 0
	StateReady   = 1
	StatePlaying = 2
	StateEnd     = 3
)

type Match struct {
	id      int
	state   int
	startAt int
	stopAt  int
}

func (m *Match) init(players []int) {

}

package missle

import (
	"math/rand"
)

type Point struct {
	X float32
	Y float32
}

type ID int64

type Player struct {
	Id       int64
	NickName string
	Avatar   string
	// game state
	IsLeft   bool
	Position Point
	Health   int
}

func MakePositionFor(isLeft bool, order int) *Point {
	if isLeft {
		return &Point{X: 1.0 / 4, Y: 0}
	} else {
		return &Point{X: 3.0 / 4, Y: 0}
	}
}

func MakeKeyPoints(count int) []*Point {
	var (
		minDX, minDY              float32
		rangeDX, rangeDY          float32
		x, y                      float32
		paddingTop, paddingBottom float32
		ny, dy                    float32
	)
	minDX = 1.0 / 8
	minDY = 1.0 / 48
	rangeDX = 1.0 / 12
	rangeDY = 1.0 / 8

	x = -minDX
	y = 1.0 / 2
	paddingTop = 1.0 / 30
	paddingBottom = 1.0 / 30
	sign := 1 // +1 - going up, -1 - going  down

	var points = make([]*Point, count)
	for i := 0; i < count; i++ {
		points[i] = &Point{X: x, Y: y}
		if i == 0 {
			x = 0
			y = 1.0 / 2
		} else {
			x = rand.Float32()*rangeDX + minDX

			for {
				dy = rand.Float32()*rangeDY + minDY
				ny = y + dy*float32(sign)
				if ny < 1-paddingTop && ny > paddingBottom {
					break
				}
			}
			y = ny
		}

		sign *= -1
	}

	return points
}

func CheckPosition(pos Point) bool {
	if pos.X < 0.0 || pos.X > 1.0 {
		return false
	}
	if pos.Y < 0.0 || pos.Y > 1.0 {
		return false
	}

	return true
}

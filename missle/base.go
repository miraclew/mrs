package missle

import (
	"fmt"
	"math/rand"
)

type Point struct {
	X float32
	Y float32
}

func (p *Point) String() string {
	return fmt.Sprintf("{%f,%f}", p.X, p.Y)
}

func pAdd(p1 Point, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func pMultiply(p Point, t float32) Point {
	return Point{p.X * t, p.Y * t}
}

type Curve struct {
	Points []*Point
}

func (c *Curve) String() string {
	var s string
	for i := 0; i < len(c.Points); i++ {
		s += fmt.Sprintf("%s \n", c.Points[i])
	}
	return s
}

func CreateBodyMovingCurve(position Point, velocity Point,
	acceleration Point, steps int, deltaTime float32) *Curve {

	curve := Curve{}
	curve.Points = make([]*Point, 0)
	for i := 0; i < steps; i++ {
		position = pAdd(position, pMultiply(velocity, deltaTime))
		velocity = pAdd(velocity, pMultiply(acceleration, deltaTime))
		curve.Points = append(curve.Points, &position)
	}
	return &curve
}

type ID int64

type Player struct {
	Id       int64
	NickName string
	Avatar   string
	// game state
	MatchId   int64
	IsLeft    bool
	Position  *Point
	Health    int32
	PointsWin int32
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
	minDX = 1.0 / 12
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
			y = 0.45
		} else {
			x += rand.Float32()*rangeDX + minDX

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

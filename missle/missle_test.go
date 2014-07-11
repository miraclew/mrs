package missle

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	// var x float32
	// x = 1.0 * 2
	// y := x * 2
	// fmt.Printf("%f, %f", x, y)
	points := MakeKeyPoints(16)
	for _, p := range points {
		fmt.Printf("%f, %f\n", p.X, p.Y)
	}

}

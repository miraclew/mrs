package main

import (
	"fmt"
	"github.com/miraclew/mrs/missle"
	"testing"
)

func TestA(t *testing.T) {
	kps := missle.MakeKeyPoints(3)
	points := generateHill(kps)

	for i := 0; i < len(points); i++ {
		fmt.Printf("p%d %f/%f \n", i, points[i].X, points[i].Y)
	}
}

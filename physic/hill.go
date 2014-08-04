package main

import (
	"github.com/miraclew/mrs/missle"
	// "log"
	"math"
)

const (
	kHillSegmentWidth = 0.005
)

func generateHill(kps []*missle.Point) []*missle.Point {
	points := make([]*missle.Point, 0)
	for i := 0; i < len(kps)-1; i++ {
		p0 := kps[i]
		p1 := kps[i+1]
		segments := int((p1.X - p0.X) / kHillSegmentWidth)
		// log.Printf("segments:%d", segments)
		dx := (p1.X - p0.X) / float32(segments)
		da := math.Pi / float32(segments)
		ymid := (p0.Y + p1.Y) / 2.0
		ampl := (p0.Y - p1.Y) / 2.0

		pt0 := p0
		pt1 := &missle.Point{}
		for j := 0; j < segments+1; j++ {
			pt1.X = p0.X + float32(j)*dx
			pt1.Y = ymid + ampl*float32(math.Cos(float64(da)*float64(j)))
			// log.Printf("p%d %f/%f \n", j, pt0.X, pt0.Y)
			points = append(points, &missle.Point{X: pt0.X, Y: pt0.Y})
			pt0 = pt1
		}
		points = append(points, &missle.Point{X: pt1.X, Y: pt1.Y})
	}

	return points
}

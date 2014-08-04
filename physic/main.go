package main

import (
	"code.google.com/p/draw2d/draw2d"
	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
	"fmt"
	"github.com/miraclew/mrs/missle"
	"image"
	// "math"
)

func drawPath(gc *draw2d.ImageGraphicContext, w int, h int, points []*missle.Point, translate bool) {
	for i := 1; i < len(points); i++ {
		gc.BeginPath()
		p0 := *points[i-1]
		p1 := *points[i]
		if translate {
			p0 = translatePoint(w, h, p0)
			p1 = translatePoint(w, h, p1)
		}
		gc.MoveTo(float64(p0.X), float64(float32(h)-p0.Y))
		gc.LineTo(float64(p1.X), float64(float32(h)-p1.Y))
		gc.Stroke()
	}
}

func translatePoint(w int, h int, p missle.Point) missle.Point {
	return missle.Point{p.X * float32(w), p.Y * float32(h)}
}

func draw(gc *draw2d.ImageGraphicContext, w int, h int) {
	fmt.Printf("w:%d h:%d", w, h)
	kps := missle.MakeKeyPoints(16)
	points := generateHill(kps)
	drawPath(gc, w, h, points, true)

	curve := missle.CreateBodyMovingCurve(missle.Point{100, 500},
		missle.Point{10, 20}, missle.Point{0, -10}, 100, 0.1)
	fmt.Println(curve)
	drawPath(gc, w, h, curve.Points, false)
}

func main() {
	window, err := x11.NewWindow()
	if err != nil {
		fmt.Printf("Cannot open an x11 window\n")
		return
	}
	screen := window.Screen()
	if rgba, ok := screen.(*image.RGBA); ok {
		gc := draw2d.NewGraphicContext(rgba)
		gc.SetStrokeColor(image.Black)
		gc.SetFillColor(image.White)
		gc.Clear()

		draw(gc, screen.Bounds().Dx(), screen.Bounds().Dy())
		fmt.Printf("This is an rgba image\n")

		window.FlushImage()

		gc.SetLineWidth(3)
		for {

			switch evt := (<-window.EventChan()).(type) {
			case ui.KeyEvent:
				if evt.Key == 'q' {
					window.Close()
				}
			}
		}
	} else {
		fmt.Printf("Not an RGBA image!\n")
	}
}

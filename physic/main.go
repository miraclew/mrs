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

func drawPath(gc *draw2d.ImageGraphicContext, w int, h int, points []*missle.Point) {
	gc.MoveTo(float64(points[0].X*float32(w)), float64(points[0].Y*float32(h)))
	for i := 1; i < len(points); i++ {
		gc.BeginPath()
		gc.MoveTo(float64(points[i-1].X*float32(w)), float64(points[i-1].Y*float32(h)))
		gc.LineTo(float64(points[i].X*float32(w)), float64(points[i].Y*float32(h)))
		gc.Stroke()
	}
}

func draw(gc *draw2d.ImageGraphicContext, w int, h int) {
	fmt.Printf("w:%d h:%d", w, h)
	kps := missle.MakeKeyPoints(16)
	points := generateHill(kps)
	drawPath(gc, w, h, points)

	curve := missle.CreateBodyMovingCurve(missle.Point{100, 100},
		missle.Point{10, 10}, missle.Point{0, -10}, 100, 0.1)
	fmt.Println(curve)
	drawPath(gc, w, h, curve.Points)
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

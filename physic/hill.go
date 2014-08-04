package main

import (
	"github.com/miraclew/mrs/missle"
	"math"
)

const (
	kHillSegmentWidth = 20
)

func generateHill(kps []*missle.Point) []*missle.Point {
	points := make([]*missle.Point, 0)
	for i := 0; i < len(kps); i++ {
		p0 := kps[i]
		p1 := kps[i+1]
		segments := int((p1.X - p0.X) / kHillSegmentWidth)
		dx := (p1.X - p0.X) / segments
		da := math.Pi / segments
		ymid := (p0.Y + p1.Y) / 2
		ampl := (p0.Y - p1.Y) / 2

		var pt0, pt1 missle.Point
		pt0 = p0
		for j := 0; j < segments+1; j++ {
			pt1.X = p0.X + j*dx
			pt1.Y = ymid + ampl*math.Cos(da*j)

			append(points, pt0)
			pt0 = pt1
		}
		append(points, pt1)
	}

	return points
}

/*

-(void)generateTerrian{
    CGMutablePathRef pathToDraw = CGPathCreateMutable();
    CGPathMoveToPoint(pathToDraw, NULL, 0.0, 0.0);
    for (int i=0; i<kMaxHillKeyPoints - 1; i++) {
        CGPoint p0 = _hillKeyPoints[i];
        CGPoint p1 = _hillKeyPoints[i+1];
        int hSegments = floorf((p1.x-p0.x)/kHillSegmentWidth);
        float dx = (p1.x - p0.x) / hSegments;
        float da = M_PI / hSegments;
        float ymid = (p0.y + p1.y) / 2;
        float ampl = (p0.y - p1.y) / 2;

        CGPoint pt0, pt1;
        pt0 = p0;
        for (int j = 0; j < hSegments+1; ++j) {

            pt1.x = p0.x + j*dx;
            pt1.y = ymid + ampl * cosf(da*j);

            CGPathAddLineToPoint(pathToDraw, NULL, pt0.x, pt0.y);
            CGPathAddLineToPoint(pathToDraw, NULL, pt1.x, pt1.y);
            pt0 = pt1;
        }

    }
//    CGPathCloseSubpath(pathToDraw);
    self.path = pathToDraw;
    [self setStrokeColor:[UIColor redColor]];
    self.physicsBody = [SKPhysicsBody bodyWithEdgeChainFromPath:pathToDraw];
    self.physicsBody.friction = 1.0f;
    self.physicsBody.usesPreciseCollisionDetection = YES;
}

*/

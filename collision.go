package packme

// collision uses SAT to compute collision between two rectangles
//
// Refer to https://gamedevelopment.tutsplus.com/tutorials/collision-detection-using-the-separating-axis-theorem--gamedev-169
func collision(i1, i2 *Item, ax, ay Axis) bool {
	dims1 := i1.Dimensions()
	dims2 := i2.Dimensions()

	// compute center of rect 1
	centerX1 := i1.pos[ax] + dims1[ax]*0.5
	centerY1 := i1.pos[ay] + dims1[ay]*0.5

	// compute center of rect 2
	centerX2 := i2.pos[ax] + dims2[ax]*0.5
	centerY2 := i2.pos[ay] + dims2[ay]*0.5

	x := max(centerX1, centerX2) - min(centerX1, centerX2)
	y := max(centerY1, centerY2) - min(centerY1, centerY2)

	return x < (dims1[ax]+dims2[ax])*0.5 &&
		y < (dims1[ay]+dims2[ay])*0.5
}

func max(x, y float32) float32 {
	if x > y {
		return x
	}

	return y
}

func min(x, y float32) float32 {
	if x < y {
		return x
	}

	return y
}

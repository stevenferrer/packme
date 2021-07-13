package packme

import "fmt"

type Point [3]float32

func NewPoint(x, y, z float32) Point {
	return Point{x, y, z}
}

func (p Point) X() float32 {
	return p[0]
}

func (p Point) Y() float32 {
	return p[1]
}

func (p Point) Z() float32 {
	return p[2]
}

func (p Point) String() string {
	return fmt.Sprintf("%v,%v,%v", p[0], p[1], p[2])
}

func computePivot(axis Axis, pos Point, dims Dimensions) Point {
	// get pivot point
	var pivot Point
	switch axis {
	case AxisX: // length-wise
		pivot = NewPoint(pos.X()+dims.Length(), pos.Y(), pos.Z())
	case AxisY: // width-wise
		pivot = NewPoint(pos.X(), pos.Y()+dims.Width(), pos.Z())
	case AxisZ: // height-wise
		pivot = NewPoint(pos.X(), pos.Y(), pos.Z()+dims.Height())
	}

	return pivot
}

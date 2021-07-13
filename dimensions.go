package packme

import "fmt"

// Dimensions is the dimensions
type Dimensions [3]float32

func NewDimensions(length, width, height float32) Dimensions {
	return Dimensions{length, width, height}
}

func (d Dimensions) Length() float32 {
	return d[0]
}

func (d Dimensions) Width() float32 {
	return d[1]
}

func (d Dimensions) Height() float32 {
	return d[2]
}

func (d Dimensions) String() string {
	return fmt.Sprintf("dims(%vx%vx%v)", d[0], d[1], d[2])
}

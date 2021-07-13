package packme

import "fmt"

type Dims [3]float32

func NewDims(length, width, height float32) Dims {
	return Dims{length, width, height}
}

func (d Dims) Length() float32 {
	return d[0]
}

func (d Dims) Width() float32 {
	return d[1]
}

func (d Dims) Height() float32 {
	return d[2]
}

func (d Dims) String() string {
	return fmt.Sprintf("dims(%vx%vx%v)", d[0], d[1], d[2])
}

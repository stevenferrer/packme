package packme

import (
	"bytes"
	"fmt"
)

type Item struct {
	desc string
	length,
	width,
	height float32
	rot Rotation
	pos Point
}

func NewItem(desc string, dims Dims) *Item {
	return &Item{
		desc:   desc,
		length: dims.Length(),
		width:  dims.Width(),
		height: dims.Height(),
		// default rotation
		rot: RotationLWH,
		// default position
		pos: Point{0, 0, 0},
	}
}

func (i *Item) Desc() string {
	return i.desc
}

func (i *Item) Length() float32 {
	return i.length
}

func (i *Item) Width() float32 {
	return i.width
}

func (i *Item) Height() float32 {
	return i.height
}

func (i *Item) Volume() float32 {
	return i.length * i.width * i.height
}

func (i *Item) Dimensions() Dims {
	var d Dims
	switch i.rot {
	case RotationLWH:
		d = Dims{i.Length(), i.Width(), i.Height()}
	case RotationWLH:
		d = Dims{i.Width(), i.Length(), i.Height()}
	case RotationWHL:
		d = Dims{i.Width(), i.Height(), i.Length()}
	case RotationHLW:
		d = Dims{i.Height(), i.Length(), i.Width()}
	case RotationHWL:
		d = Dims{i.Height(), i.Width(), i.Length()}
	case RotationLHW:
		d = Dims{i.Length(), i.Height(), i.Width()}
	}

	return d
}

func (i *Item) Collision(i2 *Item) bool {
	return collision(i, i2, AxisX, AxisY) &&
		collision(i, i2, AxisY, AxisZ) &&
		collision(i, i2, AxisX, AxisZ)
}

func (i *Item) String() string {
	return fmt.Sprintf("%s %s vol(%v) pos(%s) %s", i.desc, i.Dimensions(),
		i.Volume(), i.pos.String(), i.rot.String())
}

type Items []*Item

func (items Items) String() string {
	buf := &bytes.Buffer{}
	for _, item := range items {
		buf.WriteString(item.String())
		buf.WriteRune('\n')
	}

	return buf.String()
}

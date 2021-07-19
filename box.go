package packme

import (
	"bytes"
	"fmt"
)

type Box struct {
	desc  string
	dims  Dimensions
	items Items
}

func NewBox(desc string, dims Dimensions) *Box {
	return &Box{
		desc:  desc,
		dims:  dims,
		items: make(Items, 0),
	}
}

func (bx *Box) Desc() string {
	return bx.desc
}

func (bx *Box) Volume() float32 {
	return bx.dims.Length() * bx.dims.Width() * bx.dims.Height()
}

func (bx *Box) Dimensions() Dimensions {
	return bx.dims
}

func (bx *Box) Items() Items {
	return bx.items[:]
}

func (bx *Box) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%s items(%v) %s vol(%v)",
		bx.Desc(), len(bx.items), bx.Dimensions(), bx.Volume()))
	buf.WriteRune('\n')
	buf.WriteString(bx.items.String())
	return buf.String()
}

type Boxes []*Box

func (bxs Boxes) String() string {
	buf := &bytes.Buffer{}
	for _, bx := range bxs {
		buf.WriteString(bx.String())
		buf.WriteRune('\n')
	}
	return buf.String()
}

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

func (bx *Box) packItem(item *Item, pivot Point) bool {
	var packed bool

	// Set item position to pivot
	item.pos = pivot
	// Try posible packing rotations
	for rot := RotationLWH; rot <= RotationLHW; rot++ {
		// Set item rotation
		item.rot = rot
		// Get item dimensions based on current rotatoin
		idims := item.Dimensions()

		// Box dimensions must be able accomodate
		// the item in current rotation
		if bx.dims.Length() < pivot.X()+idims.Length() ||
			bx.dims.Width() < pivot.Y()+idims.Width() ||
			bx.dims.Height() < pivot.Z()+idims.Height() {
			continue
		}
		// Assume item is packed
		packed = true

		// See if current item might collide with other
		// items inside the box
		for _, i := range bx.items {
			if i.Collision(item) {
				// Can't pack coz it collides with another item
				packed = false
				break
			}
		}

		// Item is packed
		if packed {
			// Add item to the box
			bx.items = append(bx.items, item)
		}

		break
	}

	if !packed {
		// Reset item position
		item.pos = NewPoint(0, 0, 0)
	}

	return packed
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

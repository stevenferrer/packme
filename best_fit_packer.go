package packme

import (
	"sort"
)

// BestFitPacker is a packer implementation based on
// a paper by Dube, E., & Kanavathy, L. (2006)
type BestFitPacker struct{}

var _ Packer = (*BestFitPacker)(nil)

func NewBestFitPacker() BestFitPacker {
	return BestFitPacker{}
}

func (bfp BestFitPacker) Pack(boxSpecs []BoxSpec, itemSpecs []ItemSpec) PackingScheme {
	// Sort boxes and items by volume
	sort.Sort(BoxSpecByVolume(boxSpecs))
	sort.Sort(ItemSpecByVolume(itemSpecs))

	// List of boxes
	boxes := []*Box{}
	// List of items not packed
	notPacked := Items{}

	for _, itemSpec := range itemSpecs {
		// Loop until we pack all items in this spec
		for itemQty := itemSpec.qty; itemQty > 0; itemQty-- {
			// Spawn new item from spec
			newItem := NewItem(itemSpec.desc, itemSpec.dims)

			// Indicates whether item was packed
			packed := false

			// Try packing item to existing boxes
			for _, box := range boxes {
				packed = bfp.pack(box, newItem)
				if packed {
					// Item was packed, we're done
					break
				}
			}

			// Item doesn't fit in the existing boxes
			if !packed {
				// Try other available boxes
				for i, bxSpec := range boxSpecs {
					// See if we still have an available box of this kind
					if bxSpec.qty > 0 {
						// Spawn new box from this spec
						box := NewBox(bxSpec.desc, bxSpec.dims)

						// See if item fits in the new box
						packed = bfp.pack(box, newItem)
						if packed {
							// Reduce number of available boxes
							boxSpecs[i].qty--
							// Add new box to existing boxes
							boxes = append(boxes, box)
							break
						}
					}
				}

				// Item could not be packed
				if !packed {
					notPacked = append(notPacked, newItem)
				}
			}
		}
	}

	return PackingScheme{Boxes: boxes, NotPacked: notPacked}
}

func (bfp BestFitPacker) pack(box *Box, newItem *Item) bool {
	// Box doesn't have any items yet
	if len(box.items) < 1 {
		return bfp.packToBox(box, newItem, NewPoint(0, 0, 0))
	}

	// Try possible placements for the new item
	for axis := AxisX; axis <= AxisZ; axis++ {
		// Try packing item next to existing items in the box
		for _, boxItem := range box.items {
			// Compute pivot point
			pivot := computePivot(axis, boxItem.pos, boxItem.Dimensions())
			// Try packing the item to the pivot point
			if bfp.packToBox(box, newItem, pivot) {
				return true
			}
		}
	}

	return false
}

func (bfp BestFitPacker) packToBox(bx *Box, item *Item, pivot Point) bool {
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

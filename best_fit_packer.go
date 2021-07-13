package packme

import (
	"sort"
)

// BestFitPacker is a packer implementation based on a paper by
// Dube, E. & Kanavathy, L. (2006) titled "Optimizing Three-Dimensional
// Bin Packing Through Simulation"
//
// Refer to https://www.researchgate.net/publication/228974015_Optimizing_Three-Dimensional_Bin_Packing_Through_Simulation
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
				packed = bfp.packToBox(box, newItem)
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
						packed = bfp.packToBox(box, newItem)
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

func (bfp BestFitPacker) packToBox(box *Box, newItem *Item) bool {
	// Box doesn't have any items yet
	if len(box.items) < 1 {
		return box.packItem(newItem, NewPoint(0, 0, 0))
	}

	// Try possible placements for the new item
	for axis := AxisX; axis <= AxisZ; axis++ {
		// Try packing item next to existing items in the box
		for _, boxItem := range box.items {
			// Compute pivot point
			pivot := computePivot(axis, boxItem.pos, boxItem.Dimensions())
			// Try packing the item to the pivot point
			if box.packItem(newItem, pivot) {
				return true
			}
		}
	}

	return false
}

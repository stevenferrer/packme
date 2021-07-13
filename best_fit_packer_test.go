package packme_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sf9v/packme"
)

func TestBestFitPacker(t *testing.T) {
	boxSpecs := []packme.BoxSpec{
		packme.NewBoxSpec("Box A", 1, packme.NewDimensions(30, 30, 30)),
		packme.NewBoxSpec("Box B", 1, packme.NewDimensions(5, 5, 40)),
		packme.NewBoxSpec("Box C", 1, packme.NewDimensions(20, 20, 30)),
	}

	itemSpecs := []packme.ItemSpec{
		packme.NewItemSpec("Item A1", 17, packme.NewDimensions(10, 10, 30)),
		packme.NewItemSpec("Item A2", 1, packme.NewDimensions(10, 10, 30)),
		packme.NewItemSpec("Tall Item", 1, packme.NewDimensions(5, 39.5, 5)),
		packme.NewItemSpec("Large Item", 1, packme.NewDimensions(50, 50, 100)),
	}

	packer := packme.NewBestFitPacker()

	packingScheme := packer.Pack(boxSpecs, itemSpecs)

	assert.Len(t, packingScheme.Boxes, 3)
	assert.Len(t, packingScheme.Boxes[0].Items(), 1)
	assert.Len(t, packingScheme.Boxes[1].Items(), 4)
	assert.Len(t, packingScheme.Boxes[2].Items(), 9)
	assert.Len(t, packingScheme.NotPacked, 6)

	t.Log(packingScheme.String())
}

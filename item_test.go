package packme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItem(t *testing.T) {
	var l, w, h float32 = 10, 10, 30
	desc := "Item 1"
	item1 := NewItem(desc, NewDimensions(l, w, h))

	assert.Equal(t, desc, item1.Desc())
	assert.Equal(t, l, item1.Length())
	assert.Equal(t, w, item1.Width())
	assert.Equal(t, h, item1.Height())
	assert.Equal(t, l*w*h, item1.Volume())

	t.Run("Dimensions by rotation", func(t *testing.T) {
		t.Run("LWH", func(t *testing.T) {
			item1.rot = RotationLWH
			expect := NewDimensions(l, w, h)
			got := item1.Dimensions()
			assert.Equal(t, expect, got)
		})

		t.Run("WLH", func(t *testing.T) {
			item1.rot = RotationWLH
			expect := NewDimensions(w, l, h)
			got := item1.Dimensions()
			assert.Equal(t, expect, got)
		})

		t.Run("WHL", func(t *testing.T) {
			item1.rot = RotationWHL
			expect := NewDimensions(w, h, l)
			got := item1.Dimensions()
			assert.Equal(t, expect, got)
		})

		t.Run("HLW", func(t *testing.T) {
			item1.rot = RotationHLW
			expect := NewDimensions(h, l, w)
			got := item1.Dimensions()
			assert.Equal(t, expect, got)
		})

		t.Run("HWL", func(t *testing.T) {
			item1.rot = RotationHWL
			expect := NewDimensions(h, w, l)
			got := item1.Dimensions()
			assert.Equal(t, expect, got)
		})

		t.Run("LHW", func(t *testing.T) {
			item1.rot = RotationLHW
			expect := NewDimensions(l, h, w)
			got := item1.Dimensions()
			assert.Equal(t, expect, got)
		})
	})

	t.Run("Intersect", func(t *testing.T) {
		item2 := NewItem("Item 2", NewDimensions(10, 10, 30))
		expect := true
		got := item1.Collision(item2)
		assert.Equal(t, expect, got)

		item3 := NewItem("Item 3", NewDimensions(10, 10, 30))
		item3.pos = NewPoint(10, 0, 0)
		expect = false
		got = item1.Collision(item3)
		assert.Equal(t, expect, got)
	})
}

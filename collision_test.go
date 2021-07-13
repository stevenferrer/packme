package packme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollision(t *testing.T) {
	i1 := NewItem("Item 1", NewDims(10, 10, 30))
	i1.pos = NewPoint(0, 0, 0)

	i2 := NewItem("Item 2", NewDims(10, 10, 30))
	i2.pos = NewPoint(0, 0, 0)

	expect := true
	got := collision(i1, i2, AxisX, AxisY)
	assert.Equal(t, expect, got, "expecting collision between item 1 and 2")

	i3 := NewItem("Item 3", NewDims(10, 10, 30))
	i3.pos = NewPoint(10, 0, 0)
	expect = false
	got = collision(i1, i3, AxisX, AxisY)
	assert.Equal(t, expect, got, "expecting no collision between item 1 and 3")
}

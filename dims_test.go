package packme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDimensions(t *testing.T) {
	var l, w, h float32 = 10, 20, 30
	dim := NewDims(l, w, h)
	assert.Equal(t, l, dim.Length())
	assert.Equal(t, w, dim.Width())
	assert.Equal(t, h, dim.Height())

	expect := "dims(10x20x30)"
	got := dim.String()
	assert.Equal(t, expect, got)
}

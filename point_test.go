package packme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	var x, y, z float32 = 10, 20, 30
	point := NewPoint(x, y, z)
	assert.Equal(t, x, point.X())
	assert.Equal(t, y, point.Y())
	assert.Equal(t, z, point.Z())

	expect := "10,20,30"
	got := point.String()
	assert.Equal(t, expect, got)
}

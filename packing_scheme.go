package packme

import (
	"bytes"
	"fmt"
)

type PackingScheme struct {
	Boxes     Boxes
	NotPacked Items
}

func (ps PackingScheme) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString(ps.Boxes.String())
	buf.WriteString(fmt.Sprintf("Not packed (%v)\n", len(ps.NotPacked)))
	buf.WriteString(ps.NotPacked.String())
	return buf.String()
}

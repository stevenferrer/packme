package packme

type BoxSpec struct {
	desc string
	qty  int
	dims Dims
}

func NewBoxSpec(desc string, qty int, dims Dims) BoxSpec {
	return BoxSpec{desc: desc, qty: qty, dims: dims}
}

func (bs BoxSpec) Volume() float32 {
	return bs.dims.Length() *
		bs.dims.Width() *
		bs.dims.Height()
}

type BoxSpecByVolume []BoxSpec

func (s BoxSpecByVolume) Len() int { return len(s) }
func (s BoxSpecByVolume) Less(i, j int) bool {
	return s[i].Volume() < s[j].Volume()
}
func (s BoxSpecByVolume) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

package packme

type ItemSpec struct {
	desc string
	qty  int
	dims Dimensions
}

func NewItemSpec(desc string, qty int, dims Dimensions) ItemSpec {
	return ItemSpec{desc: desc, qty: qty, dims: dims}
}

func (i ItemSpec) Volume() float32 {
	return i.dims.Length() *
		i.dims.Width() *
		i.dims.Height()
}

type ItemSpecByVolume []ItemSpec

func (s ItemSpecByVolume) Len() int { return len(s) }
func (s ItemSpecByVolume) Less(i, j int) bool {
	return s[i].Volume() < s[j].Volume()
}
func (s ItemSpecByVolume) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

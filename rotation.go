package packme

type Rotation int

const (
	// RotationLWH (l, w, h)
	RotationLWH Rotation = iota + 1
	// RotationWLH (w, l, h)
	RotationWLH
	// RotationWHL (w, h, l)
	RotationWHL
	// RotationHLW (h, l, w)
	RotationHLW
	// RotationHWL (h, w, l)
	RotationHWL
	// RotationLHW (l, h, w)
	RotationLHW
)

func (r Rotation) String() string {
	return [...]string{
		"rot(l,w,h)",
		"rot(w,l,h)",
		"rot(w,h,l)",
		"rot(h,l,w)",
		"rot(h,w,l)",
		"rot(l,h,w)",
	}[r-1]
}

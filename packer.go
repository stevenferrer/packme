package packme

// Packer is an interface that wraps the pack method
type Packer interface {
	// Pack takes a box and item specs and returns a packing scheme
	Pack([]BoxSpec, []ItemSpec) PackingScheme
}

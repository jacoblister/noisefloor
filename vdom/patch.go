package vdom

// PatchType enumerated type
type PatchType int

// Patch types
const (
	Insert PatchType = iota
	Remove
	Replace
	AttrSet
	AttrRemove
	TextSet
)

// Patch is a DOM patch
type Patch struct {
	Type    PatchType
	Path    []int
	Element Element
	Attr    Attr
}

// PatchList is a series of DOM patches
type PatchList struct {
	SVGNamespace bool
	Patch        []Patch
}

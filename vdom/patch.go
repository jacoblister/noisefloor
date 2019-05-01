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
)

// Patch a DOM patch
type Patch struct {
	Type         PatchType
	SVGNamespace bool
	Path         []int
	Element      Element
	Attr         Attr
}

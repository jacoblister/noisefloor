package vdom

// PatchType enumerated type
type PatchType int

// Patch types
const (
	Insert PatchType = iota
	Remove
	Replace
	SetAttr
)

// Patch a DOM patch
type Patch struct {
	Type    PatchType
	Path    []int
	element Element
	attr    Attr
}

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

//JSON returns a JSON friendly encoding of the patch
func (p *Patch) JSON() interface{} {
	result := map[string]interface{}{}
	result["type"] = "Replace"
	return result
}

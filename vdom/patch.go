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
	Element Element
	Attr    Attr
}

//JSON returns a JSON friendly encoding of the patch
func (p *Patch) JSON() interface{} {
	result := map[string]interface{}{}
	result["type"] = "Replace"
	return result
}

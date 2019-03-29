package vdom

// dom is the current vdom as written to the dom
var dom = Element{Type: Root}

// RenderToDom recursively renders a VDOM tree into the real DOM
func RenderToDom(element Element) {
	patch := Patch{Type: Replace, Path: []int{}, Element: element}

	applyPatchToDom(&patch)
	dom = element
}

// fullDomPatch returns a patch to fully populate the DOM
func fullDomPatch() *Patch {
	patch := Patch{Type: Replace, Path: []int{}, Element: dom}
	return &patch
}

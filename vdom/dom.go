package vdom

// dom is the current vdom as written to the dom
var dom = Element{Type: Root}

// RenderToDom recursively renders a VDOM tree into the real DOM
func RenderToDom(element Element) {
	patch := Patch{Type: Replace, Path: []int{}, element: element}
	applyPatchToDom(&patch)

	dom = element
}

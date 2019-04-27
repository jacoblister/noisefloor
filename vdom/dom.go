package vdom

// dom is the current vdom as written to the dom
var dom = Element{Type: Root}

// rootComponent is the root of the component tree
var rootComponent Component

var svgNamespace bool

//SetSVGNamespace set the DOM namespace to SVG (default is HTML)
func SetSVGNamespace() {
	svgNamespace = true
}

// RenderComponentToDom renders a VDOM component
func RenderComponentToDom(component Component) {
	rootComponent = component
	RenderToDom(component.Render())
}

// UpdateComponent handles a state change in a component
func UpdateComponent(component Component) {
	RenderToDom(rootComponent.Render())
}

// RenderToDom recursively renders a VDOM tree into the real DOM
func RenderToDom(element Element) {
	patch := Patch{Type: Replace, SVGNamespace: svgNamespace, Path: []int{}, Element: element}

	applyPatchToDom(&patch)
	dom = element
}

// fullDomPatch returns a patch to fully populate the DOM
func fullDomPatch() *Patch {
	patch := Patch{Type: Replace, SVGNamespace: svgNamespace, Path: []int{}, Element: dom}
	return &patch
}

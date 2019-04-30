package vdom

var dom = Element{Type: Root}         // dom is the current vdom as written to the dom
var svgNamespace bool                 // svgNamespace indicates an SVG vs HTML document
var rootComponent Component           // rootComponent is the root of the component tree
var elementMap map[Component]*Element // elementMap links active components to elements

//SetSVGNamespace set the DOM namespace to SVG (default is HTML)
func SetSVGNamespace() {
	svgNamespace = true
}

// RenderComponentToDom renders a VDOM component
func RenderComponentToDom(component Component) {
	rootComponent = component
	elementMap = make(map[Component]*Element)

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

// addComponentMap adds to new component to the current element map during build up
func addComponentMap(component Component, element *Element) {
	elementMap[component] = element
}

// fullDomPatch returns a patch to fully populate the DOM
func fullDomPatch() *Patch {
	patch := Patch{Type: Replace, SVGNamespace: svgNamespace, Path: []int{}, Element: dom}
	return &patch
}

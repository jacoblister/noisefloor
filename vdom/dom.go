package vdom

var dom = Element{Type: Root}           // dom is the current vdom as written to the dom
var svgNamespace bool                   // svgNamespace indicates an SVG vs HTML document
var rootComponent Component             // rootComponent is the root of the component tree
var componentMap map[Component]*Element // elementMap links active components to elements

//SetSVGNamespace set the DOM namespace to SVG (default is HTML)
func SetSVGNamespace() {
	svgNamespace = true
}

// RenderComponentToDom renders a VDOM component
func RenderComponentToDom(component Component) {
	rootComponent = component
	rootElement := component.Render()

	updateDomTreeRecursive(&rootElement, []int{})
	RenderToDom(&rootElement)
}

// UpdateComponent handles a state change in a component
func UpdateComponent(component Component) {
	rootElement := component.Render()
	RenderToDom(&rootElement)
}

// RenderToDom recursively renders a VDOM tree into the real DOM
func RenderToDom(element *Element) {
	componentMap = map[Component]*Element{}

	patch := Patch{Type: Replace, SVGNamespace: svgNamespace, Path: []int{}, Element: *element}
	applyPatchToDom(&patch)

	dom = *element
}

// updateDomTreeRecursive updates the dom element path and componenet map for the whole tree
func updateDomTreeRecursive(element *Element, path []int) {
	element.Path = path
	if element.Component != nil {
		componentMap[element.Component] = element
	}

	for i := 0; i < len(element.Children); i++ {
		childPath := append(path, i)
		updateDomTreeRecursive(&element.Children[i], childPath)
	}
}

// fullDomPatch returns a patch to fully populate the DOM
func fullDomPatch() *Patch {
	patch := Patch{Type: Replace, SVGNamespace: svgNamespace, Path: []int{}, Element: dom}
	return &patch
}

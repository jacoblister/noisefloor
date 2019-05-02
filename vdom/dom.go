package vdom

var dom = Element{Type: Root}           // dom is the current vdom as written to the dom
var svgNamespace bool                   // svgNamespace indicates an SVG vs HTML document
var rootComponent Component             // rootComponent is the root of the component tree
var componentMap map[Component]*Element // elementMap links active components to elements
var componentUpdate chan Component      // componentUpdate is the background component update channel

//SetSVGNamespace set the DOM namespace to SVG (default is HTML)
func SetSVGNamespace() {
	svgNamespace = true
}

// SetDomRootElement sets the root DOM element
func SetDomRootElement(element *Element) {
	dom = *element
}

// RenderComponentToDom renders a VDOM component
func RenderComponentToDom(component Component) {
	rootComponent = component
	rootElement := component.Render()

	componentMap = map[Component]*Element{}
	updateDomTreeRecursive(&rootElement, []int{})
	SetDomRootElement(&rootElement)
}

// UpdateComponent is called when a state change in a component occurs
func UpdateComponent(component Component) {
	// rootElement := component.Render()
	// RenderToDom(&rootElement)
}

// UpdateComponentBackground allows a background process to
// notifiy a state change in a component
func UpdateComponentBackground(component Component) {
	componentUpdate <- component
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

// updateDomBegin notifies a DOM update cycle is starting
func updateDomBegin() {
}

// updateDomEnd notifies a DOM update cycle has ended,
// and returns a patch of DOM changes for the update cycle
func updateDomEnd() *Patch {
	dom = rootComponent.Render()
	patch := Patch{Type: Replace, SVGNamespace: svgNamespace, Path: []int{}, Element: dom}
	return &patch
}

// fullDomPatch returns a patch to fully populate the DOM
func fullDomPatch() *Patch {
	patch := Patch{Type: Replace, SVGNamespace: svgNamespace, Path: []int{}, Element: dom}
	return &patch
}

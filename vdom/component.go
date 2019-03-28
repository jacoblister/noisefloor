package vdom

// Component is a interface for a VDOM component
type Component interface {
	Render() Element
}

// RenderComponentToDom renders a VDOM component
func RenderComponentToDom(c Component) {
	RenderToDom(c.Render())
}

// UpdateComponent handles a state change in a component
func UpdateComponent(c Component) {
	RenderToDom(c.Render())
}

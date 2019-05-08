package vdom

// Component is a interface for a VDOM component
type Component interface {
	Render() Element
}

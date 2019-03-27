package vdom

// Component is a interface for a modular XML document part
type Component interface {
	Render() Element
}

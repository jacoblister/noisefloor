package vdom

import (
	"fmt"
)

// ElementType enumerated type
type ElementType int

// XML element types
const (
	Root ElementType = iota
	Normal
	Text
)

// Attr is an xml attribute
type Attr struct {
	Name  string
	Value interface{}
}

// Element is an xml element
type Element struct {
	Type  ElementType
	Name  string
	Attrs []Attr

	children      []Element
	eventHandlers []EventHandler
}

// MakeRootElement creates a VDOM root element
func MakeRootElement() Element {
	return Element{Type: Root}
}

// MakeElement creates an element with optional children and attributes
func MakeElement(name string, args ...interface{}) Element {
	element := Element{Type: Normal, Name: name}

	for i := 0; i < len(args); i++ {
		switch arg := args[i].(type) {
		case string:
			element.Attrs = append(element.Attrs, Attr{Name: arg, Value: args[i+1]})
			i++
		case Attr:
			if len(arg.Name) > 0 {
				element.Attrs = append(element.Attrs, arg)
			}
		case Element:
			element.children = append(element.children, arg)
		case EventHandler:
			element.eventHandlers = append(element.eventHandlers, arg)
		}
	}
	return element
}

// MakeTextElement creates a text element with specified text
func MakeTextElement(text string) Element {
	return Element{Type: Text, Attrs: []Attr{{Name: "Text", Value: text}}}
}

// AppendChild appends a child elements to the element
func (e *Element) AppendChild(child Element) {
	e.children = append(e.children, child)
}

// AttrMap returns this element's attributes as a map
// of attribute name to attribute value
func (e *Element) AttrMap() map[string]string {
	m := map[string]string{}
	for _, attr := range e.Attrs {
		m[attr.Name] = fmt.Sprintf("%v", attr.Value)
	}
	return m
}

// Compare non-recursively compares e to other. It does not check
// the child nodes since they can be a Node with any underlying type.
// If you want to compare the parent and children fields, use CompareNodes.
func (e *Element) Compare(other *Element, compareAttrs bool) (bool, string) {
	if e.Name != other.Name {
		return false, fmt.Sprintf("e.Name was %s but other.Name was %s", e.Name, other.Name)
	}
	if !compareAttrs {
		return true, ""
	}
	attrs := e.Attrs
	otherAttrs := other.Attrs
	if len(attrs) != len(otherAttrs) {
		return false, fmt.Sprintf("n has %d attrs but other has %d attrs.", len(attrs), len(otherAttrs))
	}
	for i, attr := range attrs {
		otherAttr := otherAttrs[i]
		if attr != otherAttr {
			return false, fmt.Sprintf("e.Attrs[%d] was %s but other.Attrs[%d] was %s", i, attr, i, otherAttr)
		}
	}
	return true, ""
}

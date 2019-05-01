package vdom

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
	Type          ElementType
	Name          string
	Attrs         map[string]interface{}
	Children      []Element
	EventHandlers []EventHandler

	Component Component // Component if element derives from a component
	Path      []int     // Element path from root of DOM tree
}

// MakeRootElement creates a VDOM root element
func MakeRootElement() Element {
	return Element{Type: Root}
}

// MakeElement creates an element with optional children and attributes
func MakeElement(name string, args ...interface{}) Element {
	element := Element{Type: Normal, Name: name, Attrs: map[string]interface{}{}, Children: []Element{}, EventHandlers: []EventHandler{}}

	for i := 0; i < len(args); i++ {
		switch arg := args[i].(type) {
		case string:
			element.Attrs[arg] = args[i+1]
			i++
		case Attr:
			if len(arg.Name) > 0 {
				element.Attrs[arg.Name] = arg.Value
			}
		case Element:
			element.Children = append(element.Children, arg)
		case EventHandler:
			element.EventHandlers = append(element.EventHandlers, arg)
		case Component:
			childElement := arg.Render()
			childElement.Component = arg
			element.Children = append(element.Children, childElement)
		case []Component:
			for j := 0; j < len(arg); j++ {
				childElement := arg[j].Render()
				childElement.Component = arg[j]
				element.Children = append(element.Children, childElement)
			}
		}
	}
	return element
}

// MakeTextElement creates a text element with specified text
func MakeTextElement(text string) Element {
	return Element{Type: Text, Attrs: map[string]interface{}{"Text": text}}
}

// AppendChild appends a child elements to the element
func (e *Element) AppendChild(child Element) {
	e.Children = append(e.Children, child)
}

// Compare non-recursively compares e to other. It does not check
// the child nodes since they can be a Node with any underlying type.
// If you want to compare the parent and children fields, use CompareNodes.
func (e *Element) Compare(other *Element, compareAttrs bool) (bool, string) {
	if e.Name != other.Name {
		return false, "mismatch names"
	}
	if !compareAttrs {
		return true, ""
	}
	attrs := e.Attrs
	otherAttrs := other.Attrs
	if len(attrs) != len(otherAttrs) {
		return false, "mismatch attrs"
	}
	// for i, attr := range attrs {
	// 	otherAttr := otherAttrs[i]
	// 	if attr != otherAttr {
	// 		return false, fmt.Sprintf("e.Attrs[%d] was %s but other.Attrs[%d] was %s", i, attr, i, otherAttr)
	// 	}
	// }
	return true, ""
}

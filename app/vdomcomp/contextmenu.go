package vdomcomp

import (
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

const menuWidth = 100

//ContextMenu is a selectable menu on right click
type ContextMenu struct {
	pickList PickList
	active   bool
}

//MakeContextMenu creates a new context menu with selection callback
func MakeContextMenu(x int, y int, items []string, active bool,
	listSelectFunc listSelectFunc) ContextMenu {
	contextMenu := ContextMenu{MakePickList(x, y, menuWidth, listItemHeight*len(items), items, "", listSelectFunc), active}
	return contextMenu
}

//Active flag getter Method
func (m *ContextMenu) Active() bool {
	return m.active
}

// SetActive sets the active state of the context menu
func (m *ContextMenu) SetActive(active bool) {
	m.active = active
}

//Render renders the ContextMenu component
func (m *ContextMenu) Render() vdom.Element {
	if !m.active {
		return vdom.MakeElement("g")
	}

	menu := vdom.MakeElement("g",
		vdom.MakeElement("rect",
			"id", "contextmenu-container",
			"stroke", "none",
			"fill", "white",
			"fill-opacity", 0,
			"x", 0,
			"y", 0,
			"width", 10000,
			"height", 10000,
			vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
				m.active = false
			}),
		),
		&m.pickList,
		vdom.MakeElement("rect",
			"pointer-events", "none",
			"id", "contextmenu",
			"stroke", "black",
			"fill", "none",
			"x", m.pickList.x,
			"y", m.pickList.y,
			"width", m.pickList.width,
			"height", m.pickList.height,
		),
	)
	return menu
}

package vdomcomp

import (
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

const menuWidth = 100
const menuItemHeight = 20

type menuSelectFunc func(menuItem string)

//ContextMenu is a selectable menu on right click
type ContextMenu struct {
	x              int
	y              int
	items          []string
	active         bool
	menuSelectFunc menuSelectFunc
	selectedIndex  int
}

//MakeContextMenu creates a new context menu with selection callback
func MakeContextMenu(x int, y int, items []string, active bool,
	menuSelectFunc menuSelectFunc) ContextMenu {
	contextMenu := ContextMenu{x, y, items, active, menuSelectFunc, 0}
	return contextMenu
}

//Render renders the ContextMenu component
func (m *ContextMenu) Render() vdom.Element {
	if !m.active {
		return vdom.MakeElement("g")
	}

	menuItems := []vdom.Element{}
	for i := 0; i < len(m.items); i++ {
		if i == m.selectedIndex {
			rect := vdom.MakeElement("rect",
				"x", m.x,
				"y", m.y+(i*menuItemHeight),
				"width", menuWidth,
				"height", menuItemHeight,
				"stroke", "none",
				"fill", "lightgrey",
			)
			menuItems = append(menuItems, rect)
		}
		item := vdom.MakeElement("text",
			"font-family", "sans-serif",
			"text-anchor", "start",
			"dominant-baseline", "central",
			"font-size", 12,
			"x", m.x+menuItemHeight/2,
			"y", m.y+(i*menuItemHeight)+(menuItemHeight/2),
			vdom.MakeTextElement(m.items[i]),
		)
		menuItems = append(menuItems, item)
	}

	menu := vdom.MakeElement("g",
		"pointer-events", "all",
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
		vdom.MakeElement("g",
			"id", "contextmenu",
			"pointer-events", "all",
			vdom.MakeElement("rect",
				"stroke", "none",
				"fill", "white",
				"x", m.x,
				"y", m.y,
				"width", menuWidth,
				"height", len(m.items)*menuItemHeight,
			),
			vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
				m.selectedIndex = event.Data["OffsetY"].(int) / menuItemHeight
			}),
			vdom.MakeEventHandler(vdom.MouseLeave, func(element *vdom.Element, event *vdom.Event) {
				m.selectedIndex = -1
			}),
			vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
				m.menuSelectFunc(m.items[m.selectedIndex])
				m.active = false
			}),
			menuItems,
			vdom.MakeElement("rect",
				"id", "contextmenu",
				"stroke", "black",
				"fill", "none",
				"x", m.x,
				"y", m.y,
				"width", menuWidth,
				"height", len(m.items)*menuItemHeight,
			),
		),
	)
	return menu
}

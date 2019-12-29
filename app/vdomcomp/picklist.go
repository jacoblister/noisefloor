package vdomcomp

import (
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

const listItemHeight = 20

type listSelectFunc func(listItem string)

//PickList is a selectable list of items
type PickList struct {
	x              int
	y              int
	width          int
	height         int
	items          []string
	listSelectFunc listSelectFunc
	selectedItem   string
	pickItem       string
}

//MakePickList creates a new context menu with selection callback
func MakePickList(x int, y int, width int, height int, items []string,
	selectedText string, listSelectFunc listSelectFunc) PickList {
	return PickList{x, y, width, height, items, listSelectFunc, selectedText, ""}
}

func makeEventHandler(eventType vdom.EventType, index int, item string, handler func(index int, item string)) vdom.EventHandler {
	return vdom.MakeEventHandler(eventType, func(element *vdom.Element, event *vdom.Event) {
		handler(index, item)
	})
}

//Render renders the ContextMenu component
func (p *PickList) Render() vdom.Element {
	listItems := []vdom.Element{}
	for i := 0; i < len(p.items); i++ {
		fill := "white"
		if p.items[i] == p.pickItem {
			fill = "lightsteelblue"
		}
		if p.items[i] == p.selectedItem {
			fill = "lightgrey"
		}
		rect := vdom.MakeElement("rect",
			"id", "list-"+p.items[i],
			"x", p.x,
			"y", p.y+(i*listItemHeight),
			"width", p.width,
			"height", listItemHeight,
			"stroke", "none",
			"fill", fill,
			makeEventHandler(vdom.Click, i, p.items[i], func(index int, item string) {
				p.pickItem = ""
				p.selectedItem = item
				p.listSelectFunc(item)
			}),
			makeEventHandler(vdom.MouseEnter, i, p.items[i], func(index int, item string) {
				p.pickItem = item
			}),
			makeEventHandler(vdom.MouseLeave, i, p.items[i], func(index int, item string) {
				p.pickItem = ""
			}),
		)
		listItems = append(listItems, rect)

		item := vdom.MakeElement("text",
			"font-family", "sans-serif",
			"text-anchor", "start",
			"dominant-baseline", "central",
			"font-size", 12,
			"x", p.x+listItemHeight/2,
			"y", p.y+(i*listItemHeight)+(listItemHeight/2),
			vdom.MakeTextElement(p.items[i]),
		)
		listItems = append(listItems, item)
	}

	pickList := vdom.MakeElement("g",
		listItems,
	)
	return pickList
}

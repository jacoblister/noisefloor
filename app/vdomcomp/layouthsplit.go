package vdomcomp

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

//LayoutHSplit vertical split container, with center resize bar
type LayoutHSplit struct {
	width           int
	height          int
	divider         int
	moving          *bool
	topComponent    vdom.Component
	bottomComponent vdom.Component
	dividerSetFunc  dividerSetFunc
}

//MakeLayoutHSplit create a new Layout vertical split componenet
func MakeLayoutHSplit(width int, height int, divider int, moving *bool,
	leftComponent vdom.Component, rightComponent vdom.Component,
	dividerSetFunc dividerSetFunc) *LayoutHSplit {
	layoutHSplit := LayoutHSplit{width, height, divider, moving,
		leftComponent, rightComponent,
		dividerSetFunc}
	return &layoutHSplit
}

//Render renders the LayoutVSplit component
func (l *LayoutHSplit) Render() vdom.Element {
	disablePointerIfMoving := vdom.Attr{}
	if *l.moving {
		disablePointerIfMoving = vdom.Attr{Name: "pointer-events", Value: "none"}
	}

	e := vdom.MakeElement("g",
		// "pointer-events", "none",
		vdom.MakeElement("rect",
			"id", "h-divider",
			"stroke", "none",
			"fill", "white",
			"x", 1,
			"y", 1,
			"width", l.width-1,
			"height", l.height-1,
			vdom.MakeEventHandler(vdom.MouseUp, func(element *vdom.Element, event *vdom.Event) {
				if *l.moving == true {
					*l.moving = false
					l.dividerSetFunc(event.Data["OffsetY"].(int))
				}
			}),
			vdom.MakeEventHandler(vdom.MouseLeave, func(element *vdom.Element, event *vdom.Event) {
				if *l.moving == true {
					*l.moving = false
					l.dividerSetFunc(event.Data["OffsetY"].(int))
				}
			}),
			vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
				if *l.moving == true {
					l.dividerSetFunc(event.Data["OffsetY"].(int))
				}
			}),
		),
		vdom.MakeElement("line",
			"id", "h-dividerline",
			"stroke", "gray",
			"x1", 0,
			"y1", l.divider,
			"x2", l.width,
			"y2", l.divider,
			"cursor", "ns-resize",
			disablePointerIfMoving,
			vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
				*l.moving = true
			}),
		),
		vdom.MakeElement("g",
			l.topComponent,
			disablePointerIfMoving,
		),
		vdom.MakeElement("g",
			"transform", "translate(0,"+strconv.Itoa(l.divider)+")",
			l.bottomComponent,
			disablePointerIfMoving,
		),
	)
	return e
}

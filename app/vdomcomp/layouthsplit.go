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
	dividerSize     int
	moving          *bool
	topComponent    vdom.Component
	bottomComponent vdom.Component
	dividerSetFunc  dividerSetFunc
}

//MakeLayoutHSplit create a new Layout vertical split componenet
func MakeLayoutHSplit(width int, height int, divider int, dividerSize int, moving *bool,
	leftComponent vdom.Component, rightComponent vdom.Component,
	dividerSetFunc dividerSetFunc) *LayoutHSplit {
	layoutHSplit := LayoutHSplit{width, height, divider, dividerSize, moving,
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
		vdom.MakeElement("rect",
			"id", "h-divider",
			"stroke", "none",
			"fill", "rgba(0,0,0,0)",
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
		vdom.MakeElement("rect",
			"id", "h-dividerline",
			"stroke", "gray",
			"fill", "gray",
			"x", 0,
			"y", l.divider,
			"width", l.width,
			"height", l.dividerSize,
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
			"transform", "translate(0,"+strconv.Itoa(l.divider+l.dividerSize)+")",
			l.bottomComponent,
			disablePointerIfMoving,
		),
	)
	return e
}

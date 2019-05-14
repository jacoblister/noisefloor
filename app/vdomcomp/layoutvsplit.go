package vdomcomp

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

type dividerSetFunc func(pos int)

//LayoutVSplit vertical split container, with center resize bar
type LayoutVSplit struct {
	width          int
	height         int
	divider        int
	moving         *bool
	leftComponent  vdom.Component
	rightComponent vdom.Component
	dividerSetFunc dividerSetFunc
}

//MakeLayoutVSplit create a new Layout horizontal split componenet
func MakeLayoutVSplit(width int, height int, divider int, moving *bool,
	leftComponent vdom.Component, rightComponent vdom.Component,
	dividerSetFunc dividerSetFunc) *LayoutVSplit {
	layoutVSplit := LayoutVSplit{width, height, divider, moving,
		leftComponent, rightComponent,
		dividerSetFunc}
	return &layoutVSplit
}

//Render renders the LayoutHSplit component
func (l *LayoutVSplit) Render() vdom.Element {
	disablePointerIfMoving := vdom.Attr{}
	if *l.moving {
		disablePointerIfMoving = vdom.Attr{Name: "pointer-events", Value: "none"}
	}

	e := vdom.MakeElement("g",
		vdom.MakeElement("rect",
			"id", "v-divider",
			"stroke", "none",
			"fill", "white",
			"x", 1,
			"y", 1,
			"width", l.width-1,
			"height", l.height-1,
			vdom.MakeEventHandler(vdom.MouseUp, func(element *vdom.Element, event *vdom.Event) {
				if *l.moving == true {
					*l.moving = false
					l.dividerSetFunc(event.Data["OffsetX"].(int))
				}
			}),
			vdom.MakeEventHandler(vdom.MouseLeave, func(element *vdom.Element, event *vdom.Event) {
				if *l.moving == true {
					*l.moving = false
					l.dividerSetFunc(event.Data["OffsetX"].(int))
				}
			}),
			vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
				if *l.moving == true {
					l.dividerSetFunc(event.Data["OffsetX"].(int))
				}
			}),
		),
		vdom.MakeElement("line",
			"id", "v-dividerline",
			"stroke", "gray",
			"x1", l.divider,
			"y1", 0,
			"x2", l.divider,
			"y2", l.height,
			"cursor", "ew-resize",
			disablePointerIfMoving,
			vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
				*l.moving = true
			}),
		),
		vdom.MakeElement("g",
			l.leftComponent,
			disablePointerIfMoving,
		),
		vdom.MakeElement("g",
			"transform", "translate("+strconv.Itoa(l.divider)+",0)",
			disablePointerIfMoving,
			l.rightComponent,
		),
	)
	return e
}

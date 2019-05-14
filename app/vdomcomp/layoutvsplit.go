package vdomcomp

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

type dividerSetFunc func(pos int)
type movingSetFunc func(moving bool)

//LayoutVSplit vertical split container, with center resize bar
type LayoutVSplit struct {
	width          int
	height         int
	divider        int
	moving         bool
	leftComponent  vdom.Component
	rightComponent vdom.Component
	dividerSetFunc dividerSetFunc
	movingSetFunc  movingSetFunc
}

//MakeLayoutVSplit create a new Layout vertical split componenet
func MakeLayoutVSplit(width int, height int, divider int, moving bool,
	leftComponent vdom.Component, rightComponent vdom.Component,
	dividerSetFunc dividerSetFunc, movingSetFunc movingSetFunc) *LayoutVSplit {
	layoutVSplit := LayoutVSplit{width, height, divider, moving,
		leftComponent, rightComponent,
		dividerSetFunc, movingSetFunc}
	return &layoutVSplit
}

//Render renders the LayoutVSplit component
func (l *LayoutVSplit) Render() vdom.Element {
	disablePointer := vdom.Attr{}
	if l.moving {
		disablePointer = vdom.Attr{Name: "pointer-events", Value: "none"}
	}

	e := vdom.MakeElement("g",
		vdom.MakeElement("rect",
			"id", "divider",
			"stroke", "gray",
			"fill", "white",
			"x", 0,
			"y", 0,
			"width", 640,
			"height", 480,
			vdom.MakeEventHandler(vdom.MouseUp, func(element *vdom.Element, event *vdom.Event) {
				if l.moving == true {
					l.moving = false
					l.dividerSetFunc(event.Data["OffsetX"].(int))
					l.movingSetFunc(l.moving)
				}
			}),
			vdom.MakeEventHandler(vdom.MouseLeave, func(element *vdom.Element, event *vdom.Event) {
				if l.moving == true {
					l.moving = false
					l.dividerSetFunc(event.Data["OffsetX"].(int))
					l.movingSetFunc(l.moving)
				}
			}),
			vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
				if l.moving == true {
					l.dividerSetFunc(event.Data["OffsetX"].(int))
				}
			}),
		),
		vdom.MakeElement("line",
			"id", "dividerline",
			"stroke", "gray",
			"x1", l.divider-1,
			"y1", 0,
			"x2", l.divider+1,
			"y2", l.height,
			disablePointer,
			"cursor", "ew-resize",
			vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
				l.moving = true
				l.movingSetFunc(l.moving)
			}),
		),
		vdom.MakeElement("g",
			l.leftComponent,
		),
		vdom.MakeElement("g",
			"transform", "translate("+strconv.Itoa(l.divider)+",0)",
			l.rightComponent,
		),
	)
	return e
}

package vdomcomp

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

//LayoutVSplit vertical split container, with center resize bar
type LayoutVSplit struct {
	width          int
	height         int
	divider        int
	leftComponent  vdom.Component
	rightComponent vdom.Component

	moving bool
}

//MakeLayoutVSplit create a new Layout vertical split componenet
func MakeLayoutVSplit(width int, height int, divider int, leftComponent vdom.Component, rightComponent vdom.Component) *LayoutVSplit {
	layoutVSplit := LayoutVSplit{width, height, divider, leftComponent, rightComponent, false}
	return &layoutVSplit
}

//Render renders the LayoutVSplit component
func (l *LayoutVSplit) Render() vdom.Element {
	e := vdom.MakeElement("g",
		vdom.MakeElement("rect",
			"id", "divider",
			"stroke", "gray",
			"fill", "white",
			"x", 0,
			"y", 0,
			"width", 640,
			"height", 480,
			vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
				l.moving = true
			}),
			vdom.MakeEventHandler(vdom.MouseUp, func(element *vdom.Element, event *vdom.Event) {
				l.moving = false
			}),
			vdom.MakeEventHandler(vdom.MouseEnter, func(element *vdom.Element, event *vdom.Event) {
				buttons := event.Data["Buttons"].(int)
				if buttons > 0 {
					l.moving = true
				}
			}),
			vdom.MakeEventHandler(vdom.MouseLeave, func(element *vdom.Element, event *vdom.Event) {
				l.moving = false
			}),
			vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
				if l.moving == true {
					l.divider = event.Data["ClientX"].(int)
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
			// "pointer-events", "none",
			"cursor", "ew-resize",
			vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
				l.moving = true
			}),
			vdom.MakeEventHandler(vdom.MouseUp, func(element *vdom.Element, event *vdom.Event) {
				l.moving = false
			}),
			vdom.MakeEventHandler(vdom.MouseEnter, func(element *vdom.Element, event *vdom.Event) {
				buttons := event.Data["Buttons"].(int)
				if buttons > 0 {
					l.moving = true
				}
			}),
			vdom.MakeEventHandler(vdom.MouseLeave, func(element *vdom.Element, event *vdom.Event) {
				l.moving = false
			}),
			vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
				if l.moving == true {
					l.divider = event.Data["ClientY"].(int)
				}
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

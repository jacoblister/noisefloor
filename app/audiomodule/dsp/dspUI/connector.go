package dspUI

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

// Connector is a connection between two Processors
type Connector struct {
	x1          int
	y1          int
	x2          int
	y2          int
	isConnected bool
	value       float32
}

// Render displays a Connector
func (c *Connector) Render() vdom.Element {
	distance := 10
	stroke := "grey"
	if c.isConnected {
		stroke = "darkblue"
	}

	path := "M" + strconv.Itoa(c.x1) + ".5 " + strconv.Itoa(c.y1) + ".5 "
	path += "L" + strconv.Itoa(c.x1+distance) + ".5 " + strconv.Itoa(c.y1) + ".5 "
	path += "L" + strconv.Itoa(c.x2-distance) + ".5 " + strconv.Itoa(c.y2) + ".5 "
	path += "L" + strconv.Itoa(c.x2) + ".5 " + strconv.Itoa(c.y2) + ".5 "

	pathElement := vdom.MakeElement("path",
		"d", path,
		"stroke", stroke,
		"fill", "none",
	)
	elem := vdom.MakeElement("g", []vdom.Element{pathElement})
	if c.isConnected {
		fill := "black"
		if c.value >= -1 && c.value < 0 {
			level := int(-c.value * 255)
			fill = "rgb(" + strconv.Itoa(level) + ",0,0)"
		} else if c.value >= 0 && c.value <= 1 {
			level := int(c.value * 255)
			fill = "rgb(0," + strconv.Itoa(level) + ",0)"
		}

		rect := vdom.MakeElement("rect",
			"x", c.x1-procConnWidth/2,
			"y", c.y1-procConnWidth/2,
			"width", procConnWidth,
			"height", procConnWidth,
			"stroke", "black",
			"fill", fill,
			"pointer-events", "none",
		)
		elem.AppendChild(rect)

		rect = vdom.MakeElement("rect",
			"x", c.x2-procConnWidth/2,
			"y", c.y2-procConnWidth/2,
			"width", procConnWidth,
			"height", procConnWidth,
			"stroke", "black",
			"fill", fill,
			"pointer-events", "none",
		)
		elem.AppendChild(rect)
	}

	return elem
}

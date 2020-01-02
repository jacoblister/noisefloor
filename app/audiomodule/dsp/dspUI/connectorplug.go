package dspUI

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

// ConnectorPlug is this signal animated connector ends
type ConnectorPlug struct {
	x1          int
	y1          int
	x2          int
	y2          int
	isConnected bool
	value       float32
}

// Render displays a Connector
func (c *ConnectorPlug) Render() vdom.Element {
	elem := vdom.MakeElement("g")
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
			"class", "animated",
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
			"class", "animated",
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

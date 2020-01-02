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
	return pathElement
	// elem := vdom.MakeElement("g", []vdom.Element{pathElement})
	//
	// return elem
}

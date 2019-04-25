package app

import (
	"github.com/jacoblister/noisefloor/vdom"
)

var nf noiseFloor

// App is the main entry point for the application
func App(driver Driver) {
	hardwareDevices := HardwareDevices{}
	driver.Start(hardwareDevices, &nf)

	vdom.RenderComponentToDom(&nf)
	vdom.ListenAndServe()
}

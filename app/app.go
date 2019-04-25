package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jacoblister/noisefloor/vdom"
)

var c components

// App is the main entry point for the application
func App(driver Driver) {
	hardwareDevices := HardwareDevices{}
	driver.Start(hardwareDevices, &c)

	go func() {
		vdom.RenderComponentToDom(&c)
		vdom.ListenAndServe()
	}()

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	driver.Stop(hardwareDevices)
}

package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jacoblister/noisefloor/audiomodule"
	"github.com/jacoblister/noisefloor/vdom"
)

var mods modules

//GetAudioProcessor returns the audioProcessor to external javascript
func GetAudioProcessor() audiomodule.AudioProcessor {
	return &mods
}

// App is the main entry point for the application
func App(driver Driver) {
	mods.synthEngine.Load("") // Load synth engine patch

	hardwareDevices := HardwareDevices{}
	driver.Start(hardwareDevices, &mods)

	go func() {
		vdom.SetSVGNamespace()
		vdom.RenderComponentToDom(&mods)
		vdom.ListenAndServe()
	}()

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	driver.Stop(hardwareDevices)
}

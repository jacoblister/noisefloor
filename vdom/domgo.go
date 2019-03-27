//+build !js

package vdom

import "fmt"

//applyPatchToDom applies the patch for the GoLang native target
func applyPatchToDom(patch *Patch) {
	fmt.Println("GoLang target apply patch", patch)
}

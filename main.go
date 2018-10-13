package main

import "fmt"
import "github.com/jacoblister/noisefloor/common"

func Process() {
	fmt.Println("do process")
}

func main() {
	// var x = 20
	var z = common.Add(1, 2)
	var event common.Midievent
	fmt.Println("hello world: %d", z)
	fmt.Println("%v", event)
}

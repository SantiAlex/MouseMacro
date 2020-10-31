package main

import (
	"flag"
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var keyButton = flag.String("k", "", "key button")
var xPosition = flag.Int("x", 0, "mouse x position")
var yPosition = flag.Int("y", 0, "mouse y position")

func main() {
	flag.Parse()
	fmt.Println("--- Please press ", *keyButton, *xPosition, *yPosition)
	robotgo.EventHook(hook.KeyDown, []string{*keyButton}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		robotgo.MoveClick(*xPosition, *yPosition, "left", true)
		robotgo.MoveMouse(x, y)
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
	"github.com/toqueteos/webbrowser"
)

func main() {
	fmt.Println("Activating scholar button")
	scholar()
}
func scholar() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(sch hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"w"}, func(sch hook.Event) {
		fmt.Println("w")

		webbrowser.Open("http://golang.org")
	})

	start := hook.Start()
	<-hook.Process(start)
}

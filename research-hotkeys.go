package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/toqueteos/webbrowser"
	"github.com/zyedidia/clipper"
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
		getstring()
		pastestring()
		webbrowser.Open("http://golang.org")
		webbrowser.Open("https://scholar.google.com.au/scholar?hl=en&as_sdt=0%2C5&q=text%20goes%20here&btnG=")
		web1 := "https://scholar.google.com.au/scholar?hl=en&as_sdt=0%2C5&q="
		web2 := "&btnG="
		web3 := web1 + web2
	})

	start := hook.Start()
	<-hook.Process(start)
}

func getstring() {
	//robotgo.TypeStr("Hello")
	robotgo.KeyTap("c", "cmd")

}
func pastestring() {
	clip, _ := clipper.GetClipboard(clipper.Clipboards...)

	// paste to stdout
	data2, _ := clip.ReadAll(clipper.RegClipboard)
	fmt.Print(string(data2))
}

/*
	func Example() {
		//clipboard.WriteAll("text")
		selected =
		text, _ := clipboard.ReadAll()
		fmt.Println(text)

		// Output:
		// 日本語
	}
*/

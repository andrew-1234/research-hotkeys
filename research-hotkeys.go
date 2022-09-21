package main

import (
	"fmt"
	"time"

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
	date := time.Now()
	fmt.Println("Press ctrl + shift + q to exit the program")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(sch hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "g"}, func(sch hook.Event) {
		getstring()
		data2 := pastestring()
		//webbrowser.Open("https://scholar.google.com.au/scholar?hl=en&as_sdt=0%2C5&q=text%20goes%20here&btnG=")
		web1 := "https://scholar.google.com.au/scholar?hl=en&as_sdt=0%2C5&q="
		web2 := "&btnG="
		web3 := web1 + string(data2) + web2
		webbrowser.Open(string(web3))
		fmt.Println(date.Format("2006-01-02"), "_scholar_", string(data2))
	})

	start := hook.Start()
	<-hook.Process(start)
}

func getstring() {
	//robotgo.TypeStr("Hello")
	robotgo.KeyTap("c", "cmd")

}

func pastestring() []byte {
	clip, _ := clipper.GetClipboard(clipper.Clipboards...)

	// paste to stdout
	data2, _ := clip.ReadAll(clipper.RegClipboard)
	return data2
	//fmt.Print(string(data2))
}

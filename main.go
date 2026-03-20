package main

import (
	_ "embed"

	"fyne.io/systray"
	"github.com/gordonklaus/portaudio"
)

//go:embed redmic.png
var redIconData []byte

//go:embed graymic.png
var grayIconData []byte

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	in := make([]int16, 1024)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), &in)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(grayIconData) // Initial gray icon
	systray.SetTitle("Mic Monitor")

	go func() {
		in := make([]int16, 1024)
		stream, _ := portaudio.OpenDefaultStream(1, 0, 44100, len(in), &in)
		stream.Start()
		defer stream.Stop()
		defer stream.Close()

		for {
			stream.Read()
			var sum int64
			for _, v := range in {
				if v < 0 {
					v = -v
				}
				sum += int64(v)
			}
			avg := sum / int64(len(in))

			if avg > 0 { // Threshold > 0
				systray.SetIcon(redIconData)
			} else {
				systray.SetIcon(grayIconData)
			}
		}
	}()
}

func onExit() {}


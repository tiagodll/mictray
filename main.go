package main

import (
	_ "embed"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"fyne.io/systray"
)

//go:embed redmic.png
var redIconData []byte

//go:embed graymic.png
var grayIconData []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(grayIconData) // Initial gray icon
	systray.SetTitle("Mic Monitor")

	var currentVol atomic.Int32

	systray.SetOnTapped(func() {
		go func() {
			if currentVol.Load() == 0 {
				exec.Command("osascript", "-e", "set volume input volume 100").Run()
				currentVol.Store(100)
				systray.SetIcon(redIconData)
			} else {
				exec.Command("osascript", "-e", "set volume input volume 0").Run()
				currentVol.Store(0)
				systray.SetIcon(grayIconData)
			}
		}()
	})

	go func() {
		for {
			cmd := exec.Command("osascript", "-e", "input volume of (get volume settings)")
			out, err := cmd.Output()
			if err == nil {
				volStr := strings.TrimSpace(string(out))
				vol, err := strconv.Atoi(volStr)
				if err == nil {
					currentVol.Store(int32(vol))
					if vol > 0 {
						systray.SetIcon(redIconData)
					} else {
						systray.SetIcon(grayIconData)
					}
				}
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()
}

func onExit() {}

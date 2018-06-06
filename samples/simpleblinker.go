// +build ignore

// Simple LED blinker, works OOTB on a RPi. However, it does not clean up
// after itself. So might leave the LED On. The RPi is not harmed though.

package main

import (
	"time"

	"github.com/weidanhome2018/embd"
	_ "github.com/weidanhome2018/embd/host/rpi" // This loads the RPi driver
)

func main() {
	for {
		embd.LEDToggle("LED0")
		time.Sleep(250 * time.Millisecond)
	}
}

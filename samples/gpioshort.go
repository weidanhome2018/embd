// +build ignore

package main

import (
	"flag"

	"github.com/weidanhome2018/embd"

	_ "github.com/weidanhome2018/embd/host/all"
)

func main() {
	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	embd.SetDirection(10, embd.Out)
	embd.DigitalWrite(10, embd.High)
}

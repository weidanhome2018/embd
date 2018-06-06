// +build ignore

package main

import (
	"flag"
	"fmt"

	"github.com/weidanhome2018/embd"

	_ "github.com/weidanhome2018/embd/host/all"
)

func main() {
	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	val, _ := embd.AnalogRead(0)
	fmt.Printf("Reading: %v\n", val)
}

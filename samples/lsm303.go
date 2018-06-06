// +build ignore

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/weidanhome2018/embd"
	"github.com/weidanhome2018/embd/sensor/lsm303"

	_ "github.com/weidanhome2018/embd/host/all"
)

func main() {
	flag.Parse()

	if err := embd.InitI2C(); err != nil {
		panic(err)
	}
	defer embd.CloseI2C()

	bus := embd.NewI2CBus(1)

	mems := lsm303.New(bus)
	defer mems.Close()

	for {
		heading, err := mems.Heading()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Heading is %v\n", heading)

		time.Sleep(500 * time.Millisecond)
	}
}

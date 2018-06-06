// +build ignore

package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/weidanhome2018/embd/controller/servoblaster"
	"github.com/weidanhome2018/embd/motion/servo"

	_ "github.com/weidanhome2018/embd/host/all"
)

func main() {
	flag.Parse()

	sb := servoblaster.New()
	defer sb.Close()

	pwm := sb.Channel(0)

	servo := servo.New(pwm)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	turnTimer := time.Tick(500 * time.Millisecond)
	left := true

	servo.SetAngle(90)
	defer func() {
		servo.SetAngle(90)
	}()

	for {
		select {
		case <-turnTimer:
			left = !left
			var err error
			switch left {
			case true:
				err = servo.SetAngle(45)
			case false:
				err = servo.SetAngle(135)
			}
			if err != nil {
				panic(err)
			}
		case <-c:
			return
		}
	}
}

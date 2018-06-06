// Package all conviniently loads all the inbuilt/supported host drivers.
package all

import (
	_ "github.com/weidanhome2018/embd/host/bbb"
	_ "github.com/weidanhome2018/embd/host/hfga"
	_ "github.com/weidanhome2018/embd/host/rpi"
)

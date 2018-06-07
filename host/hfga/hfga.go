/*
	Package rpi provides Raspberry Pi (including A+/B+) support.
	The following features are supported on Linux kernel 3.8+

	GPIO (digital (rw))
	I²C
	LED
*/
package hfga

import (
	"github.com/weidanhome2018/embd"
	"github.com/weidanhome2018/embd/host/generic"
)

var spiDeviceMinor = 0

var pins = embd.PinMap{
	&embd.PinDesc{ID: "GPIO_93", Aliases: []string{"93", "GPIO_93", "PowerButton"}, Caps: embd.CapDigital, DigitalLogical: 93},
	&embd.PinDesc{ID: "GPIO_110", Aliases: []string{"110", "GPIO_110", "key_dis"}, Caps: embd.CapDigital , DigitalLogical: 110},
	&embd.PinDesc{ID: "GPIO_148", Aliases: []string{"148", "GPIO_148", "key_wake"}, Caps: embd.CapDigital, DigitalLogical: 148},
}

var ledMap = embd.LEDMap{
	"redled": []string{"0", "redled", "REDLED"},
}

func init() {
	embd.Register(embd.HostHFGAIMX6, func(rev int) *embd.Descriptor {
		return &embd.Descriptor{
			GPIODriver: func() embd.GPIODriver {
				return embd.NewGPIODriver(pins, generic.NewDigitalPin, nil, nil)
			},
			I2CDriver: func() embd.I2CDriver {
				return embd.NewI2CDriver(generic.NewI2CBus)
			},
			LEDDriver: func() embd.LEDDriver {
				return embd.NewLEDDriver(ledMap, generic.NewLED)
			},
			SPIDriver: func() embd.SPIDriver {
				return embd.NewSPIDriver(spiDeviceMinor, generic.NewSPIBus, nil)
			},
		}
	})
}

package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {

	r := raspi.NewAdaptor()
	sensor := gpio.NewPIRMotionDriver(r, "8", 1*time.Millisecond)
	led := gpio.NewLedDriver(r, "7")

	work := func() {

		sensor.On(gpio.MotionDetected, func(data interface{}) {
			fmt.Println("Motion is detected by sensor! ")
			led.On()
		})

		sensor.On(gpio.MotionStopped, func(data interface{}) {
			fmt.Println("Sensor stopped!")
			led.Off()
		})
	}

	robot := gobot.NewRobot("R3MotionBot",
		[]gobot.Connection{r},
		[]gobot.Device{sensor, led},
		work,
	)

	robot.Start()
}

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
	led := gpio.NewLedDriver(r, "7")

	err := led.Off()

	if err == nil {
		work := func() {

			fmt.Println("Running led test....")
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("R3AppBot",
			[]gobot.Connection{r},
			[]gobot.Device{led},
			work,
		)
		robot.Start()
	} else {
		fmt.Println("Error with the LED....{}", err)
	}

}

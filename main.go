package main

import (
	"flag"
	"main/components/dbmanager"
	"main/components/gpiodriver"
	"main/components/inputprocessor"
	"main/components/reviewprocessor"
	"main/components/server"
	"os"
	"os/signal"
	"time"

	"github.com/pterm/pterm"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	//display banner
	banner()

	//Parse customization flags
	dbPtr := flag.String("db", "analytics.db", "name for database file")
	portPtr := flag.Int("port", 8080, "port to listen on")
	logPtr := flag.Bool("log", false, "log to stdout")
	pin1Ptr := flag.Int("pin1", 17, "pin for 1 star rating")
	pin2Ptr := flag.Int("pin2", 27, "pin for 2 star rating")
	pin3Ptr := flag.Int("pin3", 22, "pin for 3 star rating")
	pin4Ptr := flag.Int("pin4", 10, "pin for 4 star rating")
	pin5Ptr := flag.Int("pin5", 9, "pin for 5 star rating")
	flag.Parse()

	//capture ctrl+c or cmd+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Info.Println("\nExiting...")

		err := dbmanager.Close()
		if err != nil {
			pterm.Error.Println(err)
		}

		err1 := gpiodriver.Close()
		if err1 != nil {
			pterm.Error.Println(err)
		}

		if err != nil || err1 != nil {
			os.Exit(1)
		}

		os.Exit(0)
	}()

	//Initialize GPIO
	InitializeGPIO(*pin1Ptr, *pin2Ptr, *pin3Ptr, *pin4Ptr, *pin5Ptr)

	//Initialize database
	InitializeDB(dbPtr)

	//Fork a process that listens for input and adds the rating to the database
	go inputprocessor.AddRatingtoDB()

	//Start server
	server.Start(*portPtr, *logPtr)
}

func InitializeGPIO(pin1ID, pin2ID, pin3ID, pin4ID, pin5ID int) {
	err := gpiodriver.Open()
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}

	//Register bcm2835 pins on the Raspberry Pi and their corresponding ratings

	errors := make([]error, 0)

	gpiodriver.RegisterInputPin(pin1ID)
	gpiodriver.RegisterInputPin(pin2ID)
	gpiodriver.RegisterInputPin(pin3ID)
	gpiodriver.RegisterInputPin(pin4ID)
	gpiodriver.RegisterInputPin(pin5ID)

	pin1, err := gpiodriver.GetPin(pin1ID)
	if err != nil {
		errors = append(errors, err)
	}
	pin1.Detect(rpio.AnyEdge)

	pin2, err := gpiodriver.GetPin(pin2ID)
	if err != nil {
		errors = append(errors, err)
	}
	pin2.Detect(rpio.AnyEdge)

	pin3, err := gpiodriver.GetPin(pin3ID)
	if err != nil {
		errors = append(errors, err)
	}
	pin3.Detect(rpio.AnyEdge)

	pin4, err := gpiodriver.GetPin(pin4ID)
	if err != nil {
		errors = append(errors, err)
	}
	pin4.Detect(rpio.AnyEdge)

	pin5, err := gpiodriver.GetPin(pin5ID)
	if err != nil {
		errors = append(errors, err)
	}
	pin5.Detect(rpio.AnyEdge)

	go func() {
		for {
			if pin1.EdgeDetected() {
				if pin1.Read() == rpio.High {
					gpiodriver.InputChan <- pin1ID
				}
			}

			if pin2.EdgeDetected() {
				if pin2.Read() == rpio.High {
					gpiodriver.InputChan <- pin2ID
				}
			}

			if pin3.EdgeDetected() {
				if pin3.Read() == rpio.High {
					gpiodriver.InputChan <- pin3ID
				}
			}

			if pin4.EdgeDetected() {
				if pin4.Read() == rpio.High {
					gpiodriver.InputChan <- pin4ID
				}
			}

			if pin5.EdgeDetected() {
				if pin5.Read() == rpio.High {
					gpiodriver.InputChan <- pin5ID
				}
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	errors = append(errors, inputprocessor.RegisterPintoRating(pin1ID, 1))
	errors = append(errors, inputprocessor.RegisterPintoRating(pin2ID, 2))
	errors = append(errors, inputprocessor.RegisterPintoRating(pin3ID, 3))
	errors = append(errors, inputprocessor.RegisterPintoRating(pin4ID, 4))
	errors = append(errors, inputprocessor.RegisterPintoRating(pin5ID, 5))

	checkforErrors(errors)
}

func InitializeDB(dbPtr *string) {
	//Initialize database
	err := dbmanager.Open(*dbPtr)
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}

	//Automatically create Ratings table if it doesn't exist
	err = dbmanager.AutoCreateStruct(&reviewprocessor.DateAnalytics{})
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
}

func checkforErrors(errors []error) {
	for _, err := range errors {
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}
	}
}

//banner displays the application banner
func banner() {
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("StarTrack"))
	pterm.Info.Println("Made by Akhil Datla")
}

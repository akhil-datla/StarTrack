package gpiodriver

import (
	"errors"
	"fmt"
	"sync"

	"github.com/pterm/pterm"
	"github.com/stianeikeland/go-rpio/v4"
)

var pins = make(map[int]rpio.Pin)
var mutex = &sync.Mutex{}
var InputChan = make(chan int)

//Open the connection to the GPIO pins
func Open() error {
	return rpio.Open()
}

//RegisterInputPin registers a pin as an input pin
func RegisterInputPin(pin int) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := pins[pin]; ok {
		pterm.Error.Print("Pin already registered")
	}
	pinObj := rpio.Pin(pin)
	pinObj.Input()
	pins[pin] = pinObj
}

//RegisterOutputPin registers a pin as an output pin
func RegisterOutputPin(pin int) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := pins[pin]; ok {
		pterm.Error.Print("Pin already registered")
	}
	pinObj := rpio.Pin(pin)
	pinObj.Output()
	pins[pin] = pinObj
}

//GetPin returns the pin object for the given pin number
func GetPin(pin int) (rpio.Pin, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := pins[pin]; !ok {
		return 0, errors.New("pin not registered")
	}
	pinObj := pins[pin]
	return pinObj, nil
}

//Write writes to the pin setting the state to high (true) or low (false)
func Write(pin int, state bool) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := pins[pin]; !ok {
		return errors.New("pin not registered")
	}
	pinObj := pins[pin]
	if state {
		pinObj.High()
	} else {
		pinObj.Low()
	}
	return nil
}

//Close the connection to the GPIO pins
func Close() error {
	return rpio.Close()
}

//GetPins returns the list of pins that have been registered
func GetPins() []string {
	pins := make([]string, 0)
	mutex.Lock()
	defer mutex.Unlock()
	for k := range pins {
		pins = append(pins, fmt.Sprintf("%d", k))
	}
	return pins
}

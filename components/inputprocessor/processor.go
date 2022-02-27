package inputprocessor

import (
	"errors"
	"main/components/gpiodriver"
	"main/components/reviewprocessor"
	"sync"

	"github.com/pterm/pterm"
)

var ratingPins = make(map[int]int)
var mutex = &sync.Mutex{}

//RegisterPintoRating registers a pin to a rating
func RegisterPintoRating(pin, rating int) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := ratingPins[pin]; ok {
		return errors.New("pin already registered")
	}

	ratingPins[pin] = rating
	return nil
}

//GetPintoRatings gets the coresponding ratings for pins
func GetPintoRatings() map[int]int {
	mutex.Lock()
	defer mutex.Unlock()
	return ratingPins
}

//AddRatingtoDB adds a rating to the database once an input is detected
func AddRatingtoDB() {
	for {
		select {
		case pin := <-gpiodriver.InputChan:
			mutex.Lock()

			if _, ok := ratingPins[pin]; !ok {
				mutex.Unlock()
				pterm.Error.Print("Rating not registed for pin")
			} else {
				rating := ratingPins[pin]

				mutex.Unlock()

				err := reviewprocessor.SaveReview(rating)
				if err != nil {
					pterm.Error.Print(err)
				}
			}

		}
	}
}

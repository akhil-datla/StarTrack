package reviewprocessor

import (
	"errors"
	"main/components/dbmanager"
	"time"

	uuid "github.com/satori/go.uuid"
)

//DateAnalytics is a struct that represents the ratings for a date
type DateAnalytics struct {
	ID      string `storm:"id" json:"id"`
	Date    string `storm:"index" json:"date"`
	Rating1 int    `json:"rating1"`
	Rating2 int    `json:"rating2"`
	Rating3 int    `json:"rating3"`
	Rating4 int    `json:"rating4"`
	Rating5 int    `json:"rating5"`
}

//SaveReview saves a review to the database
func SaveReview(rating int) error {
	date := time.Now().Format("01-02-2006")
	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	dateAnalytics := DateAnalytics{}
	err := dbmanager.Query("Date", date, &dateAnalytics)

	if err != nil {
		if err.Error() == "not found" {
			if rating == 1 {
				dateAnalytics.Rating1++
			}
			if rating == 2 {
				dateAnalytics.Rating2++
			}
			if rating == 3 {
				dateAnalytics.Rating3++
			}
			if rating == 4 {
				dateAnalytics.Rating4++
			}
			if rating == 5 {
				dateAnalytics.Rating5++
			}

			dateAnalytics.ID = uuid.NewV4().String()
			dateAnalytics.Date = date

			err := dbmanager.Save(&dateAnalytics)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if rating == 1 {
			dateAnalytics.Rating1++
		}
		if rating == 2 {
			dateAnalytics.Rating2++
		}
		if rating == 3 {
			dateAnalytics.Rating3++
		}
		if rating == 4 {
			dateAnalytics.Rating4++
		}
		if rating == 5 {
			dateAnalytics.Rating5++
		}

		err := dbmanager.Update(&dateAnalytics)
		if err != nil {
			return err
		}
	}
	return nil
}

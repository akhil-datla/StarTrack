package reportgenerator

import (
	"main/components/dbmanager"
	"main/components/reviewprocessor"
)

//GetRatings gets the total ratings for a given date
func GetRatings(date string) (map[string]int, error) {
	ratings := make(map[string]int)
	var dateAnalytics reviewprocessor.DateAnalytics
	err := dbmanager.Query("Date", date, &dateAnalytics)

	if err != nil {
		return nil, err
	} else {
		ratings["1"] = dateAnalytics.Rating1
		ratings["2"] = dateAnalytics.Rating2
		ratings["3"] = dateAnalytics.Rating3
		ratings["4"] = dateAnalytics.Rating4
		ratings["5"] = dateAnalytics.Rating5
	}
	return ratings, nil
}

package reportgenerator

import (
	"main/components/dbmanager"
	"main/components/reviewprocessor"
	"time"
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

//GetRatingsRange gets the total ratings for a given date range
func GetRatingsRange(startDate string, endDate string) (map[string]int, error) {
	totalRatings := make(map[string]int)
	var dateAnalytics reviewprocessor.DateAnalytics

	currDateObj, err := time.Parse("01-02-2006", startDate)
	if err != nil {
		return nil, err
	}
	endDateObj, err := time.Parse("01-02-2006", endDate)
	if err != nil {
		return nil, err
	}

	for currDateObj.Before(endDateObj) {
		err = dbmanager.Query("Date", currDateObj.Format("01-02-2006"), &dateAnalytics)
		if err != nil {
			currDateObj = currDateObj.AddDate(0, 0, 1)
			continue
		} else {
			totalRatings["1"] += dateAnalytics.Rating1
			totalRatings["2"] += dateAnalytics.Rating2
			totalRatings["3"] += dateAnalytics.Rating3
			totalRatings["4"] += dateAnalytics.Rating4
			totalRatings["5"] += dateAnalytics.Rating5
		}
		currDateObj = currDateObj.AddDate(0, 0, 1)
	}
	return totalRatings, nil
}
package reportgenerator

import (
	"main/components/dbmanager"
	"main/components/reviewprocessor"
	"time"
)

//GetRatings gets the total ratings in the database
func GetRatings() (map[string]int, error) {
	totalRatings := make(map[string]int)
	var dateAnalytics []*reviewprocessor.DateAnalytics
	err := dbmanager.QueryAll(&dateAnalytics)
	if err != nil {
		totalRatings["1"] = 0
		totalRatings["2"] = 0
		totalRatings["3"] = 0
		totalRatings["4"] = 0
		totalRatings["5"] = 0
		return totalRatings, err
	}

	for _, dateAnalytic := range dateAnalytics {
		totalRatings["1"] += dateAnalytic.Rating1
		totalRatings["2"] += dateAnalytic.Rating2
		totalRatings["3"] += dateAnalytic.Rating3
		totalRatings["4"] += dateAnalytic.Rating4
		totalRatings["5"] += dateAnalytic.Rating5
	}

	return totalRatings, nil

}

//GetRatingsDate gets the total ratings for a given date
func GetRatingsDate(date string) (map[string]int, error) {
	ratings := make(map[string]int)
	var dateAnalytics reviewprocessor.DateAnalytics
	err := dbmanager.Query("Date", date, &dateAnalytics)

	if err != nil {
		ratings["1"] = 0
		ratings["2"] = 0
		ratings["3"] = 0
		ratings["4"] = 0
		ratings["5"] = 0
		return ratings, err
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

	totalRatings["1"] = 0
	totalRatings["2"] = 0
	totalRatings["3"] = 0
	totalRatings["4"] = 0
	totalRatings["5"] = 0

	currDateObj, err := time.Parse("01-02-2006", startDate)
	if err != nil {
		return totalRatings, err
	}
	endDateObj, err := time.Parse("01-02-2006", endDate)
	if err != nil {
		return totalRatings, err
	}
	newEndDateObj := endDateObj.AddDate(0, 0, 1)

	for currDateObj.Before(newEndDateObj) {
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
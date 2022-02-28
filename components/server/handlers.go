package server

import (
	"main/components/reportgenerator"
	"net/http"

	"github.com/labstack/echo/v4"
)

//getRatings is an API handler that gets the total ratings in the database
func getRatings(c echo.Context) error {
	ratings, err := reportgenerator.GetRatings()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ratings)
}

//getRatingsDate is an API handler that gets the total ratings for a given date
func getRatingsDate(c echo.Context) error {
	jsonMap := make(map[string]interface{})
	err := c.Bind(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	date := jsonMap["date"].(string) //date is in the format mm-dd-yyyy
	ratings, err := reportgenerator.GetRatingsDate(date)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ratings)
}

//getRatingsRange is an API handler that gets the total ratings for a given date range
func getRatingsRange(c echo.Context) error {
	jsonMap := make(map[string]interface{})
	err := c.Bind(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	startDate := jsonMap["startDate"].(string) //startDate is in the format mm-dd-yyyy
	endDate := jsonMap["endDate"].(string)    //endDate is in the format mm-dd-yyyy
	ratings, err := reportgenerator.GetRatingsRange(startDate, endDate)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ratings)
}

package server

import (
	"main/components/reportgenerator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getRatings(c echo.Context) error {
	date := c.QueryParam("date") //Should be in this format: 02-26-2022
	ratings, err := reportgenerator.GetRatings(date)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ratings)
}
package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

//Start starts the server
func Start(port int, log bool) {
	e = echo.New()
	e.HideBanner = true
	e.Use(middleware.Gzip())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "build",
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))
	InitializeRoutes(log)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

//InitializeRoutes initializes the routes
func InitializeRoutes(log bool) {
	if log {
		e.Use(middleware.Logger())
	}

	e.Use(middleware.Recover())

	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}

	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))

	e.GET("/ratings", getRatings)
	e.POST("/ratings/date", getRatingsDate)
	e.POST("/ratings/range", getRatingsRange)

}

package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

//Start starts the server
func Start(port int, log bool) {
	e = echo.New()
	e.HideBanner = true
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

	e.Static("/", filepath.Join(filepath.Dir(""), "static"))

	e.GET("/ratings", getRatings)

}

package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const (
	webDir      = "public/"
	defaultPort = "8088"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/hello", helloPage)
	e.Static("/", webDir)

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = defaultPort
	}
	listenPort = ":" + listenPort

	e.Logger.Fatal(e.Start(listenPort))
	e.Logger.Infof("listen, url: http://localhost%s/", listenPort)
}

func helloPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func customHTTPErrorHandler(err error, c echo.Context) {
	errorPage := webDir + "index.html"
	if err := c.File(errorPage); err != nil {
		//c.Logger().Error(err)
	}
	//c.Logger().Error(err)
}

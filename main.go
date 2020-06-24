package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	webDir     = "web/dist/"
	listenPort = ":8088"
)

func main() {
	e := echo.New()
	// e.HideBanner = true
	e.GET("/hello", helloPage)
	e.Static("/", webDir)
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Infof("listen, url: http://localhost%s/", listenPort)
	e.Logger.Fatal(e.Start(listenPort))
}

func helloPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func customHTTPErrorHandler(err error, c echo.Context) {
	errorPage := webDir + "index.html"
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

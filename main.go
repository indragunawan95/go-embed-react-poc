package main

import (
	"log"
	"net/http"

	"go-embed-react-poc/web"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	web.RegisterHandler(e)

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World duh")
	})

	err := e.Start(":8080")
	if err != nil {
		log.Fatal("faield to run")
	}

}

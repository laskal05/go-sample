package main

import (
	"github.com/labstack/echo"
	"github.com/laskal05/go-sample/handler"
)

func main() {
	e := echo.New()

	// e.Use(middleware.Recover())

	e.POST("/post", handler.Post())

	e.Start(":1323")
}

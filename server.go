package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Logger.Fatal(e.Start(":8000"))
}

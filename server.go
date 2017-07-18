package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

var (
	prod bool
)

func main() {
	flag.BoolVar(&prod, "prod", true, "Set to true in production to use let's encrypt")
	flag.Parse()

	e := echo.New()
	e.HideBanner = true

	if !prod {
		e.Logger.Fatal(e.Start(":8000"))
	} else {
		// redirect http requests to https
		e.Pre(middleware.HTTPSRedirect())
		go func(c *echo.Echo) {
			e.Logger.Fatal(e.Start(":80"))
		}(e)

		// if production, request certificate from let's encrypt
		// thanks mozilla!
		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("tonycai.me")
		e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
			Code: 301,
		}))

		e.Logger.Fatal(e.StartAutoTLS(":443"))
	}
}

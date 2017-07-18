package main

import (
	"flag"

	"github.com/caiyeon/lunch-with-us/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

var (
	prod bool
)

func main() {
	// command line args
	flag.BoolVar(&prod, "prod", true, "Set to true in production to use let's encrypt")
	flag.Parse()

	// initialize echo webserver
	e := echo.New()
	e.HideBanner = true

	// middleware

	// if production, add extra security measures
	if prod {
		// redirect http requests to https
		e.Pre(middleware.HTTPSRedirect())
		go func(c *echo.Echo) {
			e.Logger.Fatal(e.Start(":80"))
		}(e)

		// thanks mozilla!
		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("tonycai.me")
		e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
			Code: 301,
		}))
	}

	// api routing
	e.GET("/v1/ping", handlers.Ping())

	// launch webserver listener
	if prod {
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		e.Logger.Fatal(e.Start(":8000"))
	}
}

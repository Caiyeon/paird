package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/caiyeon/lunch-with-us/handlers"
	"github.com/caiyeon/lunch-with-us/store"
	"github.com/caiyeon/lunch-with-us/vault"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

var (
	prod bool
)

func main() {
	// command line args
	flag.BoolVar(&prod, "prod", false, "Set to true in production to use let's encrypt")
	flag.Parse()

	// setup vault token
	if vault.VaultToken = os.Getenv("VAULT_TOKEN"); vault.VaultToken == "" {
		panic("VAULT_TOKEN env var is not set!")
	}

	// setup persistence layer
	if err := store.Initialize("bolt.db"); err != nil {
		panic(err)
	}
	defer store.CloseDB()

	// initialize echo webserver
	e := echo.New()
	e.HideBanner = true

	// middleware
	e.Use(middleware.BodyDump(
		func(c echo.Context, reqBody, resBody []byte) {
			fmt.Printf("Request body:\n%s\n\n", reqBody)
			fmt.Printf("Response body:\n%s\n\n", resBody)
		}),
	)

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

	// serve static folder
	e.Static("/", "lunch-with-us/docs")

	// api routing
	e.GET("/v1/ping", handlers.Ping())
	e.POST("/v1/signup", handlers.Signup())
	e.POST("/v1/interactive", handlers.Interactive())

	// launch webserver listener
	fmt.Println("Starting webserver at port 8000")

	if prod {
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		e.Logger.Fatal(e.Start(":8000"))
	}
}

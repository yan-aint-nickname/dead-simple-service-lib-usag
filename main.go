package main

import (
	"fmt"
	"github.com/yan-aint-nickname/dead-simple-service-library-usage/some_domain"

	"github.com/gofiber/fiber/v2"
	undead "github.com/yan-aint-nickname/dead-simple-service-library"
)


func createApp(settings undead.Settings) *fiber.App {

	app := undead.CreateDefaultApp(settings)
	
	// NOTE: routes registry
	v1 := app.Group("/api/v1")
	some_domain.SomeDomainRoutes(&v1)

	return app
}

func main() {
	// TODO: Add web app settings struct initialization
	settings, _ := undead.LoadSettings()
	fmt.Println("LOOK:", settings.AppName, settings.Addr)

	app := createApp(settings)
	undead.RunFromSettings(app, settings)
}

package some_domain


import (
	"github.com/gofiber/fiber/v2"
)

func SomeDomainRoutes(group *fiber.Router) {
	domainGroup := (*group).Group("/domain")

	domainGroup.Get("/", HandleRoot)
	domainGroup.Get("/event/:name", HandleEvent)
}

package middleware

import (
	"context"

	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/DanielChachagua/GestionCar/pkg/key"
	"github.com/gofiber/fiber/v2"
)



func InjectApp(app *dependencies.Application) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext() 
		ctx = context.WithValue(ctx, key.AppKey, app) 
		c.SetUserContext(ctx)

		return c.Next()
	}
}

// func InjectTenantDB() fiber.Handler {
//     return func(c *fiber.Ctx) error {
//         tenantURI := c.Get("X-Tenant-Uri") // O como determines el tenant
//         db, err := database.GetTenantDB(tenantURI)
//         if err != nil {
//             return c.Status(500).SendString("DB error")
//         }
//         ctx := context.WithValue(c.UserContext(), key.TenantDBKey, db)
//         c.SetUserContext(ctx)
//         return c.Next()
//     }
// }
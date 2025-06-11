package middleware

import (
	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/DanielChachagua/GestionCar/pkg/key"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token no proporcionado",
			})
		}

		ctx := c.UserContext()
		deps := ctx.Value(key.AppKey).(*dependencies.Application)

		claims, err := utils.VerifyToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Claims inválidos",
			})
		}

		userId, ok := mapClaims["id"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "ID inválido en el token",
			})
		}

		isAdmin := false
		if adm, ok := mapClaims["is_admin_tenant"].(bool); ok {
			isAdmin = adm
		}

		var tenantID string
		if tid, ok := mapClaims["tenant_id"].(string); ok {
			tenantID = tid
		}

		if tenantID != "" {
			tenant, err := deps.TenantController.TenantService.TenantGetByID(tenantID)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: err.Error(),
				})
			}

			connection, err := utils.Decrypt(tenant.Connection)
			if err != nil {
				return c.Status(500).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: err.Error(),
				})
			}

			db, err := database.GetTenantDB(connection)
			if err != nil {
				return c.Status(500).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: err.Error(),
				})
			}

			ctx := c.UserContext()
			depsTenant := ctx.Value(key.TenantDBKey).(*dependencies.TenantApplication)
			depsTenant.SetDBTenantRepository(db)

			userFromToken := models.AuthenticatedUser{}
			if !isAdmin {
				user, err := depsTenant.MemberController.MemberService.MemberGetByID(userId)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
						Status:  false,
						Body:    nil,
						Message: "Error interno",
					})
				}

				permissions := func(perms []models.Permission) []string {
					names := make([]string, 0, len(perms))
					for _, p := range perms {
						names = append(names, p.Code)
					}
					return names
				}(user.Role.Permissions)

				userFromToken = models.AuthenticatedUser{
					ID:            user.ID,
					FirstName:     user.FirstName,
					LastName:      user.LastName,
					Username:      user.Username,
					IsAdminTenant: false,
					RoleID:        &user.Role.ID,
					RoleName:      &user.Role.Name,
					Permissions:   permissions,
					TenantID:      &tenantID,
					TenantName:    &tenant.Name,
					Identifier:    &tenant.Identifier,
				}
			} else {
				user, err := deps.AuthController.AuthService.CurrentUser(userId)

				if err != nil {
					if errResp, ok := err.(*models.ErrorStruc); ok {
						return c.Status(errResp.StatusCode).JSON(models.Response{
							Status:  false,
							Body:    nil,
							Message: errResp.Message,
						})
					}
					return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
						Status:  false,
						Body:    nil,
						Message: "Error interno",
					})
				}

				for _, userTenant := range user.UserTenants {
					if userTenant.TenantID == tenantID {
						userFromToken = models.AuthenticatedUser{
							ID:            user.ID,
							FirstName:     user.FirstName,
							LastName:      user.LastName,
							Username:      user.Username,
							IsAdminTenant: userTenant.IsAdmin,
							RoleID:        nil,
							RoleName:      nil,
							Permissions:   nil,
							TenantID:      &tenantID,
							TenantName:    &tenant.Name,
							Identifier:    &tenant.Identifier,
						}
						break
					}
				}
			}

			c.Locals("user", &userFromToken)

			return c.Next()
		}

		user, err := deps.AuthController.AuthService.CurrentUser(userId)

		if err != nil {
			if errResp, ok := err.(*models.ErrorStruc); ok {
				return c.Status(errResp.StatusCode).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: errResp.Message,
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "Error interno",
			})
		}

		userFromToken := models.AuthenticatedUser{
			ID:            user.ID,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Username:      user.Username,
			IsAdminTenant: true,
			RoleID:        nil,
			RoleName:      nil,
			Permissions:   nil,
			TenantID:      nil,
			TenantName:    nil,
			Identifier:    nil,
		}

		c.Locals("user", &userFromToken)

		return c.Next()
	}
}

package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"gorm.io/gorm"
)

func (r *Repository) AuthLogin(username string, password string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "User not found", err)
		}
		return nil, models.ErrorResponse(500, "Error retrieving user", err)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, models.ErrorResponse(401, "Incorrect credentials", nil)
	}

	return &user, nil
}

func (r *Repository) AuthTenant(userID string, tenantID string) (*models.Tenant, error) {
	var tenant models.Tenant
	err := r.DB.
		Model(&models.Tenant{}).
		Preload("UserTenants", "user_id = ? AND tenant_id = ?", userID, tenantID).
		Where("id = ?", tenantID).
		Scan(&tenant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Tenant not found", err)
		}
		return nil, models.ErrorResponse(500, "Error retrieving tenant", err)
	}

	if !tenant.IsActive {
		return nil, models.ErrorResponse(403, "Tenant is inactive", nil)
	}

	if len(tenant.UserTenants) == 0 {
		return nil, models.ErrorResponse(403, "You do not have permission to access this tenant", nil)
	}

	if !tenant.UserTenants[0].IsActive {
		return nil, models.ErrorResponse(403, "You do not have permission to access this tenant", nil)

	}

	return &tenant, nil
}

func (r *Repository) CurrentUser(userID string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "User not found", err)
		}
		return nil, models.ErrorResponse(500, "Error retrieving user", err)
	}
	return &user, nil
}

func (r *Repository) GetUserRolePermissions(connection, userID string) (*models.Member, *models.Role, []models.Permission, error) {
    db, err := database.GetTenantDB(connection)
    if err != nil {
        return nil, nil, nil, models.ErrorResponse(500, "Error retrieving the database", err)
    }

    var member models.Member
    if err := db.Where("user_id = ?", userID).First(&member).Error; err != nil {
        return nil, nil, nil, models.ErrorResponse(404, "User not found", err)
    }

    var role models.Role
    if err := db.Where("id = ?", member.RoleID).First(&role).Error; err != nil {
        return nil, nil, nil, models.ErrorResponse(404, "Role not found", err)
    }

    var permissions []models.Permission
    if err := db.Model(&role).Association("Permissions").Find(&permissions); err != nil {
        return nil, nil, nil, models.ErrorResponse(500, "Error al obtener los permisos", err)
    }

    return &member, &role, permissions, nil
}
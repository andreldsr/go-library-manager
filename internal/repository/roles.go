package repository

import (
	"go-library-manager/internal/database"
	"go-library-manager/internal/models"
)

func FindRolesByName(names []string) (roles []models.Role) {
	database.DB.Model(&models.Role{}).Where("name in ?", names).Scan(&roles)
	return
}

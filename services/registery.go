package services

import (
	"golang_fiber_api/services/user"

	"gorm.io/gorm"
)

// ServiceRegistry holds all initialized services
type ServiceRegistry struct {
	UserService *user.UserService
	// Add other services here
	// ProductService *ProductService
	// AuthService    *AuthService
}

// NewServiceRegistry initializes all services with DB
func NewServiceRegistry(db *gorm.DB) *ServiceRegistry {
	return &ServiceRegistry{
		UserService: user.NewUserService(db),
		// ProductService: NewProductService(db),
		// AuthService:    NewAuthService(db),
	}
}

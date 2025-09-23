package services

import "gorm.io/gorm"

// ServiceRegistry holds all initialized services
type ServiceRegistry struct {
	UserService *UserService
	// Add other services here
	// ProductService *ProductService
	// AuthService    *AuthService
}

// NewServiceRegistry initializes all services with DB
func NewServiceRegistry(db *gorm.DB) *ServiceRegistry {
	return &ServiceRegistry{
		UserService: NewUserService(db),
		// ProductService: NewProductService(db),
		// AuthService:    NewAuthService(db),
	}
}

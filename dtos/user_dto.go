package dtos

import (
	"golang_fiber_api/models"
	"time"
)

type CreateUserDTO struct {
	Name     string `json:"name" form:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type UpdateUserDTO struct {
	Name     *string `json:"name,omitempty" form:"name" validate:"omitempty,min=2,max=100"`
	Email    *string `json:"email,omitempty" form:"email" validate:"omitempty,email"`
	Password *string `json:"password,omitempty" form:"password" validate:"omitempty,min=6"`
}

type UserResponseDTO struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func ToUserResponseDTO(user models.User) UserResponseDTO {
	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	return UserResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

package dtos

type CreateUserDTO struct {
	Name     string `json:"name" form:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type UpdateUserDTO struct {
	Name     *string `json:"name,omitempty" form:"name"`
	Email    *string `json:"email,omitempty" form:"email"`
	Password *string `json:"password,omitempty" form:"password"`
}

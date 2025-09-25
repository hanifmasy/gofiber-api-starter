package user

import (
	"errors"
	"golang_fiber_api/database"
	"golang_fiber_api/dtos"
	"golang_fiber_api/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUsers(page, limit int) (dtos.PaginatedUsersResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	var users []models.User
	var totalRows int64

	s.db.Model(&models.User{}).Count(&totalRows)

	offset := (page - 1) * limit

	if err := s.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return dtos.PaginatedUsersResponse{}, err
	}

	// map to DTO
	var userDTOs []dtos.UserResponseDTO
	for _, u := range users {
		userDTOs = append(userDTOs, dtos.ToUserResponseDTO(u))
	}

	totalPages := int((totalRows + int64(limit) - 1) / int64(limit))

	return dtos.PaginatedUsersResponse{
		Meta: dtos.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalRows:  totalRows,
			TotalPages: totalPages,
		},
		Users: userDTOs,
	}, nil
}

func (UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

// Fetch user by ID (returns *models.User, not DTO)
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

// Update user in DB (returns only error)
func (s *UserService) UpdateUser(user *models.User) error {
	if err := s.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) CreateUser(dto dtos.CreateUserDTO) (*dtos.UserResponseDTO, error) {
	user := models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password, // hash in real case
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	resp := dtos.ToUserResponseDTO(user)
	return &resp, nil
}

func (s *UserService) DeleteUser(id int) (*dtos.UserResponseDTO, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return nil, err
	}

	// fetch with DeletedAt
	if err := s.db.Unscoped().First(&user, id).Error; err != nil {
		return nil, err
	}

	resp := dtos.ToUserResponseDTO(user)
	return &resp, nil
}

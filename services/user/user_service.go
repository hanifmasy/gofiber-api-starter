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

func (s *UserService) GetUsers(dtRequest *dtos.DataTableRequest) (dtos.PaginatedUsersResponse, error) {
	var users []models.User
	var totalRows int64

	query := s.db.Model(&models.User{})

	query = query.Where("deleted_at IS NULL")

	if dtRequest.Search != "" {
		searchTerm := "%" + dtRequest.Search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ?", searchTerm, searchTerm)
	}

	// Apply filters
	for key, value := range dtRequest.Filter {
		switch key {
		case "status":
			if boolVal, ok := value.(bool); ok {
				query = query.Where("status = ?", boolVal)
			}
		case "startDate":
			if dateStr, ok := value.(string); ok {
				query = query.Where("created_at >= ?", dateStr)
			}
		case "endDate":
			if dateStr, ok := value.(string); ok {
				query = query.Where("created_at <= ?", dateStr)
			}
		case "array":
			if intSlice, ok := value.([]int); ok && len(intSlice) > 0 {
				query = query.Where("some_field IN ?", intSlice)
			}
			// Add more filter cases as needed
		}
	}

	// Get total count with filters applied
	if err := query.Count(&totalRows).Error; err != nil {
		return dtos.PaginatedUsersResponse{}, err
	}

	// Apply sorting, offset, and limit
	if err := query.
		Order(dtRequest.GetOrderClause()).
		Offset(dtRequest.GetOffset()).
		Limit(dtRequest.Limit).
		Find(&users).Error; err != nil {
		return dtos.PaginatedUsersResponse{}, err
	}

	// Map to DTO
	var userDTOs []dtos.UserResponseDTO
	for _, u := range users {
		userDTOs = append(userDTOs, dtos.ToUserResponseDTO(u))
	}

	// Calculate total pages
	totalPages := int((totalRows + int64(dtRequest.Limit) - 1) / int64(dtRequest.Limit))

	if userDTOs == nil {
		userDTOs = []dtos.UserResponseDTO{}
	}

	return dtos.PaginatedUsersResponse{
		Meta: dtos.PaginationMeta{
			Page:       dtRequest.Page,
			Limit:      dtRequest.Limit,
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

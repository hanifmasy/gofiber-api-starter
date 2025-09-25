package user

import (
	"errors"
	"golang_fiber_api/database"
	"golang_fiber_api/dtos"
	"golang_fiber_api/models"
	"os"
	"time"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// Hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// Compare password
func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Signup
func (s *AuthService) Signup(dto dtos.UserSignupDTO) (*dtos.UserResponseDTO, error) {
	// validate unique email
	var existing models.User
	if err := s.db.Where("email = ?", dto.Email).First(&existing).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := hashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashed,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	resp := dtos.ToUserResponseDTO(user)
	return &resp, nil
}

// Signin
func (s *AuthService) Signin(dto dtos.UserSigninDTO) (string, error) {
	var user models.User
	if err := s.db.Where("email = ?", dto.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	if !checkPasswordHash(dto.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	// create token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Signout (stateless)
func (s *AuthService) Signout() bool {
	// in stateless JWT, signout is handled client-side (remove token)
	return true
}

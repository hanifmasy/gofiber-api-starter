package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents user model
type User struct {
	// @Description User Model
	// gorm.Model        // Includes ID, CreatedAt, UpdatedAt, DeletedAt
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Name      string         `gorm:"type:varchar(100)" json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `gorm:"type:varchar(255)" json:"-"`
}

// Model for user

package models

import (
	"time"
)

type User struct {
	// Table for user data.
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
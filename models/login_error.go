package models

import (
	"time"
)

// LoginError represents a login error record in the database
type LoginError struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"index"`
	UserName  string    `json:"user_name"`
	IP        string    `json:"ip"`
	ErrorType string    `json:"error_type"`
	ErrorMsg  string    `json:"error_msg"`
	Platform  string    `json:"platform"`
	Device    string    `json:"device"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
}

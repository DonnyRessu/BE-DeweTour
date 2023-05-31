package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullName" gorm:"type: VARCHAR(255)"`
	Email     string    `json:"email" binding:"required, email" gorm:"unique; not null"`
	Password  string    `json:"password" gorm:"type: VARCHAR(255)"`
	Phone     string    `json:"phone" gorm:"type: VARCHAR(255)"`
	Address   string    `json:"address" gorm:"type: VARCHAR(255)"`
	Role      string    `json:"role" gorm:"type: VARCHAR(25)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

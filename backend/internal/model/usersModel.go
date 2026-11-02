package model

import time "time"

type Users struct {
	UserID   int       `json:"user_id" db:"user_id"`
	Username string    `json:"username" db:"username"`
	Email	string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	AuthUid   string    `json:"auth_uid" db:"auth_uid"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Address string `json:"address" db:"address"`
	Points int `json:"points" db:"points"`
}
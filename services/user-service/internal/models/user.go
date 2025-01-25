package models

import "time"

type User struct {
	Username          string    `json:"username" db:"username"`
	HashedPassword    string    `json:"-" db:"hashed_password"`
	FullName          string    `json:"full_name" db:"full_name"`
	Email             string    `json:"email" db:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at" db:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}

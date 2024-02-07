package user

import "time"

type User struct {
	ID           int32
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

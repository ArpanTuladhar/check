package user

import "time"

type UserRepo interface {
}

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time //meta feild to see when timestamp
	UpdatedAt time.Time
}

package middleware

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

var (
	UsernameMinLength = 2
	PasswordMinLength = 6
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$") // for email regection taken from Stackoverflow

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
}

type AuthResponse struct {
	AccessToken string
	User        User
}

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)

	in.Username = strings.TrimSpace(in.Username)
}

func (in RegisterInput) Validate() error { // returning error - but pointer not used beacause no need to modify this function

	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: Username not long enough, (%d) characters at least", ErrValidation, UsernameMinLength)
	}

	if !emailRegexp.MatchString(in.Email) {
		return fmt.Errorf("%w:email not valid", ErrValidation)
	}

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: Password not long enough, (%d) characters as least", ErrValidation, PasswordMinLength)
	}
	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: confirm password must math the password", ErrValidation)
	}

	return nil
}

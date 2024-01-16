package register

import (
	"context"
	"fmt"
	"strings"

	"github.com/88labs/andpad-engineer-training/2023/Prashant/X/backend/internal/domain/model/user"
	"github.com/88labs/andpad-engineer-training/2023/Prashant/X/backend/internal/errors"
	"github.com/go-playground/validator/v10"
)

var (
	UsernameMinLength = 2
	PasswordMinLength = 6
)

var validate *validator.Validate

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
}

type AuthResponse struct {
	AccessToken string
	User        user.User
}

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

func init() {
	validate = validator.New()
}

func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)

	in.Username = strings.TrimSpace(in.Username)
}

func (in RegisterInput) Validate() error {
	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: Username not long enough, (%d) characters at least", errors.ErrValidation, UsernameMinLength)
	}

	err := validate.Var(in.Email, "required,email")
	if err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidation, err.Error())
	}

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: Password not long enough, (%d) characters as least", errors.ErrValidation, PasswordMinLength)
	}
	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: confirm password must match the password", errors.ErrValidation)
	}

	return nil
}

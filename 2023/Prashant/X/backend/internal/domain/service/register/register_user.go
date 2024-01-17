package register

import (
	"context"
	"fmt"
	"strings"

	"github.com/88labs/andpad-engineer-training/2023/Prashant/X/backend/internal/domain/model/user"
	"github.com/88labs/andpad-engineer-training/2023/Prashant/X/backend/internal/errors"
	"github.com/go-playground/validator/v10"
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
	Email           string `validate:"required,email"`
	Username        string `validate:"min=2"`
	Password        string `validate:"min=6"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
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
	if err := validate.Struct(in); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidation, err.Error())
	}
	return nil
}

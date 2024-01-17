package register

import (
	"testing"

	"github.com/88labs/andpad-engineer-training/2023/Prashant/X/backend/internal/errors"
	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Sanitize(t *testing.T) {
	input := RegisterInput{
		Username:        "harsh",
		Email:           "HARSH@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	expected := RegisterInput{
		Username:        "harsh",
		Email:           "harsh@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}
	input.Sanitize()

	require.Equal(t, expected, input)
}

func TestRegisterInput_Validate(t *testing.T) {
	TestCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "harsh",
				Email:           "harsh@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		{
			name: "Invalid Email",
			input: RegisterInput{
				Username:        "harsh",
				Email:           "harsh",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: errors.ErrValidation,
		},
		{
			name: "Very Short UserName",
			input: RegisterInput{
				Username:        "h",
				Email:           "harsh@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: errors.ErrValidation,
		},
		{
			name: "Too Short password",
			input: RegisterInput{
				Username:        "harsh",
				Email:           "harsh@gmail.com",
				Password:        "pass",
				ConfirmPassword: "pass",
			},
			err: errors.ErrValidation,
		},
		{
			name: "Confirm Password does not match",
			input: RegisterInput{
				Username:        "harsh",
				Email:           "harsh@gmail.com",
				Password:        "123password",
				ConfirmPassword: "galat password",
			},
			err: errors.ErrValidation,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

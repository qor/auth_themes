package clean

import (
	"errors"

	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"github.com/qor/auth/providers/password"
)

// ErrPasswordConfirmationNotMatch password confirmation not match error
var ErrPasswordConfirmationNotMatch = errors.New("password confirmation doesn't match password")

// New initialize clean theme
func New(config *auth.Config) *auth.Auth {
	if config == nil {
		config = &auth.Config{}
	}
	config.ViewPaths = append(config.ViewPaths, "github.com/qor/auth_themes/clean/views")

	Auth := auth.New(config)

	Auth.RegisterProvider(password.New(&password.Config{
		Confirmable: true,
		RegisterHandler: func(context *auth.Context) (*claims.Claims, error) {
			context.Request.ParseForm()

			if context.Request.Form.Get("confirm_password") != context.Request.Form.Get("password") {
				return nil, ErrPasswordConfirmationNotMatch
			}

			return password.DefaultRegisterHandler(context)
		},
	}))

	return Auth
}

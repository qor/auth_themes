package clean

import (
	"errors"

	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"github.com/qor/auth/database"
)

// ErrPasswordConfirmationNotMatch password confirmation not match error
var ErrPasswordConfirmationNotMatch = errors.New("password confirmation doesn't match password")

// New initialize clean them
func New(Auth *auth.Auth) {
	Auth.RegisterProvider(database.New(&database.Config{
		RegisterHandler: func(context *auth.Context) (*claims.Claims, error) {
			context.Request.ParseForm()

			if context.Request.Form.Get("confirm_password") != context.Request.Form.Get("password") {
				return nil, ErrPasswordConfirmationNotMatch
			}

			return database.DefaultRegisterHandler(context)
		},
	}))

	Auth.Render.PrependViewPath("github.com/qor/auth_themes/clean/views")
}

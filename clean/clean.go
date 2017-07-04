package clean

import (
	"github.com/qor/auth"
	"github.com/qor/auth/database"
)

// New initialize clean them
func New(auth *auth.Auth) {
	auth.RegisterProvider(database.New(nil))

	auth.Render.RegisterViewPath("github.com/qor/auth_themes/clean/views")
}

package devise

import (
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
)

// New initialize devise theme
func New(config *auth.Config) *auth.Auth {
	return clean.New(config)
}

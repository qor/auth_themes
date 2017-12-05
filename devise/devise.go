package devise

import (
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
)

// New initialize devise theme
func New(config *auth.Config) *auth.Auth {
	if config == nil {
		config = &auth.Config{}
	}

	return clean.New(config)
}

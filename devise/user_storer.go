package devise

import (
	"reflect"

	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/claims"
	"github.com/qor/qor/utils"
)

type UserStorer struct {
	auth.UserStorer
}

// Get defined how to get user with user id
func (UserStorer) Get(Claims *claims.Claims, context *Context) (user interface{}, err error) {
	var tx = context.Auth.GetDB(context.Request)

	if context.Auth.Config.UserModel != nil {
		if Claims.UserID != "" {
			currentUser := reflect.New(utils.ModelType(context.Auth.Config.UserModel)).Interface()
			if err = tx.First(currentUser, Claims.UserID).Error; err == nil {
				return currentUser, nil
			}
			return nil, ErrInvalidAccount
		}
	}

	var (
		authIdentity = reflect.New(utils.ModelType(context.Auth.Config.AuthIdentityModel)).Interface()
		authInfo     = auth_identity.Basic{
			Provider: Claims.Provider,
			UID:      Claims.Id,
		}
	)

	if !tx.Where(authInfo).First(authIdentity).RecordNotFound() {
		if context.Auth.Config.UserModel != nil {
			if authBasicInfo, ok := authIdentity.(interface {
				ToClaims() *claims.Claims
			}); ok {
				currentUser := reflect.New(utils.ModelType(context.Auth.Config.UserModel)).Interface()
				if err = tx.First(currentUser, authBasicInfo.ToClaims().UserID).Error; err == nil {
					return currentUser, nil
				}
				return nil, ErrInvalidAccount
			}
		}

		return authIdentity, nil
	}

	return nil, ErrInvalidAccount
}

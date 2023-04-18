package user

import (
	"context"

	"github.com/aruji1211/myGram/module/model/user"
)

type UserRepository interface {
	Create(ctx context.Context, userIn user.User) (err error)
	// Authenticate(ctx context.Context, userIn user.User) (userOut user.User, err error)
	GetByUsername(ctx context.Context, username string) (userOut user.User, err error)
}

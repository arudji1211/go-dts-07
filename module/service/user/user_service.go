package user

import (
	"github.com/aruji1211/myGram/module/model/token"
	ModelUser "github.com/aruji1211/myGram/module/model/user"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(ctx *gin.Context, userIn ModelUser.UserCreateRequest) (UserOut ModelUser.UserCreateResponse, err error)
	AuthenticateUser(ctx *gin.Context, userIn ModelUser.UserAuthenticate) (tokenOut token.Tokens, err error)
}

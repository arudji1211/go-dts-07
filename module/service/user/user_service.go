package user

import (
	"github.com/arudji1211/go-dts-07/module/model/token"
	ModelUser "github.com/arudji1211/go-dts-07/module/model/user"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(ctx *gin.Context, userIn ModelUser.UserCreateRequest) (UserOut ModelUser.UserCreateResponse, err error)
	AuthenticateUser(ctx *gin.Context, userIn ModelUser.UserAuthenticate) (tokenOut token.Tokens, err error)
}

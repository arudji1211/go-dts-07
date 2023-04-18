package user

import (
	hdluser "github.com/aruji1211/myGram/module/handle/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(v1 *gin.RouterGroup, UserHdl hdluser.UserHandler) {
	g := v1.Group("/user")
	g.POST("/register", UserHdl.Register)
	g.POST("/login", UserHdl.Login)
}

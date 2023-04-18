package socialmedia

import (
	hdlsocmed "github.com/aruji1211/myGram/module/handle/socialmedia"
	"github.com/aruji1211/myGram/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SocialmediaRoute(v1 *gin.RouterGroup, hdlSoc hdlsocmed.SocialmediaHandler) {
	s := v1.Group("/socialmedia")
	s.Use(middleware.JWTMiddleware())
	s.GET("/getall", hdlSoc.GetAll)
	s.GET("/getone/:id", hdlSoc.GetById)
	s.POST("/createsocialmedia", hdlSoc.Create)
	s.PUT("/updatesocialmedia/:id", hdlSoc.Update)
	s.DELETE("/deleteSocialmedia/:id", hdlSoc.Delete)

}

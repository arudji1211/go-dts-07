package product

import (
	hdlproduct "github.com/arudji1211/go-dts-07/module/handle/product"
	"github.com/arudji1211/go-dts-07/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoute(v1 *gin.RouterGroup, hdlSoc hdlproduct.ProductHandler) {
	s := v1.Group("/product")
	s.Use(middleware.JWTMiddleware())
	s.GET("/getone/:id", middleware.AuthorizationById(), hdlSoc.GetById)
	s.POST("/createproduct", hdlSoc.Create)
	s.PUT("/updateproduct/:id", middleware.AuthorizationByRole(), hdlSoc.Update)
	s.DELETE("/deleteproduct/:id", middleware.AuthorizationByRole(), hdlSoc.Delete)

}

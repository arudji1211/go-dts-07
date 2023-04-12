package book

import (
	"github.com/arudji1211/go-dts-07/module/controller"
	"github.com/gin-gonic/gin"
)

func BookRouter(v1 *gin.RouterGroup, bc controller.BookController) {

	g := v1.Group("/book")
	g.GET("", bc.GetAllBook)
	g.GET("/:id", bc.GetBookById)
	g.DELETE("/:id", bc.DeleteBook)
	g.POST("", bc.CreateBook)
	g.PUT("/:id", bc.UpdateBook)
}

package router

import (
	"github.com/arudji1211/go-dts-07/module/controller"
	"github.com/gin-gonic/gin"
)

func BookRouter(GEngine *gin.Engine, bc controller.BookController) *gin.Engine {
	GEngine.GET("/book", bc.GetAllBook)
	GEngine.GET("/book/:id", bc.GetBookById)
	GEngine.DELETE("/book/:id", bc.DeleteBook)
	GEngine.POST("/book", bc.CreateBook)
	GEngine.PUT("/book/:id", bc.UpdateBook)
	return GEngine
}

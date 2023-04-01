package server

import (
	"github.com/arudji1211/go-dts-07/module/controller"
	"github.com/arudji1211/go-dts-07/module/repository"
	"github.com/arudji1211/go-dts-07/module/router"
	"github.com/arudji1211/go-dts-07/module/service"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	ginServer := gin.Default()

	br := repository.NewBookRepository()
	bs := service.NewBookService(br)
	bc := controller.NewBookController(bs)

	ginServer = router.BookRouter(ginServer, bc)

	ginServer.Run(":9090")
}

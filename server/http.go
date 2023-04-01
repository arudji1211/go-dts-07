package server

import (
	"fmt"

	"github.com/arudji1211/go-dts-07/config"
	"github.com/arudji1211/go-dts-07/docs"
	"github.com/arudji1211/go-dts-07/module/router/v1/book"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHttpServer() {
	hdls := initDI()

	// init server
	ginServer := gin.Default()

	if config.Load.Server.Env == config.ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// init middleware
	ginServer.Use(
		gin.Logger(),   // untuk log request yang masuk
		gin.Recovery(), // untuk auto restart kalau panic
	)
	//swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//mendaftarkan route
	v1 := ginServer.Group("/api/v1")
	book.BookRouter(v1, hdls)

	///run server
	ginServer.Run(fmt.Sprintf(":%v", config.Load.Server.Http.Port))
}
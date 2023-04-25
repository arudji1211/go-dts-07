package server

import (
	"fmt"

	"github.com/arudji1211/go-dts-07/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/arudji1211/go-dts-07/config"
	"github.com/gin-gonic/gin"

	modelProduct "github.com/arudji1211/go-dts-07/module/model/product"
	modelUser "github.com/arudji1211/go-dts-07/module/model/user"

	"github.com/go-playground/validator/v10"

	repositoryProduct "github.com/arudji1211/go-dts-07/module/repository/product"
	repositoryUser "github.com/arudji1211/go-dts-07/module/repository/user"

	serviceProduct "github.com/arudji1211/go-dts-07/module/service/product"
	serviceUser "github.com/arudji1211/go-dts-07/module/service/user"

	handlerProduct "github.com/arudji1211/go-dts-07/module/handle/product"
	handlerUser "github.com/arudji1211/go-dts-07/module/handle/user"

	routerProduct "github.com/arudji1211/go-dts-07/module/router/v1/product"
	routerUser "github.com/arudji1211/go-dts-07/module/router/v1/user"
)

func Serve() {
	var validate *validator.Validate
	//Load Model
	pgConn := config.NewPostgresGormConn()
	MProduct := modelProduct.Product{}
	MUser := modelUser.User{}

	if config.Load.DataSource.Migrate {
		pgConn.AutoMigrate(&MUser, &MProduct)
	}
	//bookService := service.NewBookService(bookRepo)
	//bookController := controller.NewBookController(bookService)

	//load Repo
	RProduct := repositoryProduct.NewProductRepository(pgConn)
	RUser := repositoryUser.NewUserRepository(pgConn)

	//load Services
	SProduct := serviceProduct.NewProductService(RProduct, validate)
	SUser := serviceUser.NewUserService(RUser, validate)

	//load Handler
	hUser := handlerUser.NewUserHandler(SUser)
	hProduct := handlerProduct.NewProductHandler(SProduct)

	ginServer := gin.Default()

	if config.Load.Server.Env == config.ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	//###init middleware
	ginServer.Use(
		gin.Logger(),   // untuk log request yang masuk
		gin.Recovery(), // untuk auto restart kalau panic
	)
	//###swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ###mendaftarkan route

	v1 := ginServer.Group("/api/v1")

	routerUser.UserRouter(v1, hUser)
	routerProduct.ProductRoute(v1, hProduct)

	///run server
	ginServer.Run(fmt.Sprintf(":%v", config.Load.Server.Http.Port))
}
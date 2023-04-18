package server

import (
	"fmt"

	"github.com/aruji1211/myGram/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/aruji1211/myGram/config"
	"github.com/gin-gonic/gin"

	modelComment "github.com/aruji1211/myGram/module/model/comment"
	modelPhoto "github.com/aruji1211/myGram/module/model/photo"
	modelSocmed "github.com/aruji1211/myGram/module/model/socialmedia"
	modelUser "github.com/aruji1211/myGram/module/model/user"

	"github.com/go-playground/validator/v10"

	repositoryComment "github.com/aruji1211/myGram/module/repository/comment"
	repositoryPhoto "github.com/aruji1211/myGram/module/repository/photo"
	repositorySocmed "github.com/aruji1211/myGram/module/repository/socialmedia"
	repositoryUser "github.com/aruji1211/myGram/module/repository/user"

	serviceComment "github.com/aruji1211/myGram/module/service/comment"
	servicePhoto "github.com/aruji1211/myGram/module/service/photo"
	serviceSocmed "github.com/aruji1211/myGram/module/service/socialmedia"
	serviceUser "github.com/aruji1211/myGram/module/service/user"

	handlerComment "github.com/aruji1211/myGram/module/handle/comment"
	handlerPhoto "github.com/aruji1211/myGram/module/handle/photo"
	handlerSocmed "github.com/aruji1211/myGram/module/handle/socialmedia"
	handlerUser "github.com/aruji1211/myGram/module/handle/user"

	routerComment "github.com/aruji1211/myGram/module/router/v1/comment"
	routerPhoto "github.com/aruji1211/myGram/module/router/v1/photo"
	routerSocmed "github.com/aruji1211/myGram/module/router/v1/socialmedia"
	routerUser "github.com/aruji1211/myGram/module/router/v1/user"
)

func Serve() {
	var validate *validator.Validate
	//Load Model
	pgConn := config.NewPostgresGormConn()
	MPhoto := modelPhoto.Photo{}
	MSocial := modelSocmed.Socialmedia{}
	MUser := modelUser.User{}
	MComment := modelComment.Comment{}

	if config.Load.DataSource.Migrate {
		pgConn.AutoMigrate(&MPhoto, &MSocial, &MUser, &MComment)
	}
	//bookService := service.NewBookService(bookRepo)
	//bookController := controller.NewBookController(bookService)

	//load Repo
	RPhoto := repositoryPhoto.NewPhotoRepository(pgConn)
	RSocial := repositorySocmed.NewSocialmediaRepository(pgConn)
	RUser := repositoryUser.NewUserRepository(pgConn)
	RComment := repositoryComment.NewCommentRepository(pgConn)

	//load Services
	SPhoto := servicePhoto.NewPhotoService(RPhoto, validate)
	SComment := serviceComment.NewCommentService(RComment, validate)
	SUser := serviceUser.NewUserService(RUser, validate)
	SSocial := serviceSocmed.NewSocialmediaService(RSocial, validate)

	//load Handler
	hUser := handlerUser.NewUserHandler(SUser)
	hSocmed := handlerSocmed.NewSocialmediaHandler(SSocial)
	hPhoto := handlerPhoto.NewPhotoHandler(SPhoto)
	hComment := handlerComment.NewCommenHandler(SComment)

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
	routerSocmed.SocialmediaRoute(v1, hSocmed)
	routerPhoto.PhotoRoute(v1, hPhoto)
	routerComment.CommentRoute(v1, hComment)

	///run server
	ginServer.Run(fmt.Sprintf(":%v", config.Load.Server.Http.Port))
}

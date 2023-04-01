package server

import (
	"github.com/arudji1211/go-dts-07/config"
	"github.com/arudji1211/go-dts-07/module/controller"
	"github.com/arudji1211/go-dts-07/module/repository"
	"github.com/arudji1211/go-dts-07/module/service"
)

func initDI() controller.BookController {
	var bookRepo repository.BookRepository

	switch config.Load.DataSource.Mode {
	case config.MODE_GORM:
		pgConn := config.NewPostgresGormConn()
		bookRepo = repository.NewBookRepositoryGorm(pgConn)
	case config.MODE_MAP:
		bookRepo = repository.NewBookRepositoryMap()
	}

	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	return bookController
}

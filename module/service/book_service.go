package service

import (
	"github.com/arudji1211/go-dts-07/module/model"
	"github.com/gin-gonic/gin"
)

type BookService interface {
	Insert(ctx *gin.Context, request model.BookCreateRequest) model.BookResponse
	Update(ctx *gin.Context, request model.BookUpdateRequest) (model.BookResponse, error)
	Delete(ctx *gin.Context, id int) error
	GetAll(ctx *gin.Context) []model.BookResponse
	GetById(ctx *gin.Context, id int) (model.BookResponse, error)
}

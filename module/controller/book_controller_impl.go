package controller

import (
	"net/http"
	"strconv"

	"github.com/arudji1211/go-dts-07/module/handler"
	"github.com/arudji1211/go-dts-07/module/model"
	"github.com/arudji1211/go-dts-07/module/service"
	"github.com/gin-gonic/gin"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bs service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bs,
	}
}

func (bc *BookControllerImpl) CreateBook(ctx *gin.Context) {
	var req model.BookCreateRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		handler.ErrorBadReq(ctx)
		return
	}
	var resp model.WebResponse
	resp.Message = "Success Create Book"
	resp.Data = bc.BookService.Insert(ctx, req)
	ctx.JSON(http.StatusCreated, resp)
}

func (bc *BookControllerImpl) GetAllBook(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	var response model.WebResponse

	data := bc.BookService.GetAll(ctx)
	response.Data = data
	response.Message = "Sukses Mengambil Data Buku"
	ctx.JSON(http.StatusOK, response)

}

func (bc *BookControllerImpl) GetBookById(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	var response model.WebResponse
	parameter := ctx.Param("id")
	var id int
	id, err := strconv.Atoi(parameter)

	if err != nil {
		handler.ErrorBadReq(ctx)
		return
	}

	data, err := bc.BookService.GetById(ctx, id)
	if err != nil {
		handler.ErrorNotFound(ctx, err.Error())
		return
	}

	response.Message = "Sukses Mengambil Data Buku"
	response.Data = data
	ctx.JSON(http.StatusOK, response)

}

func (bc *BookControllerImpl) UpdateBook(ctx *gin.Context) {
	var req model.BookUpdateRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		handler.ErrorBadReq(ctx)
		return
	}

	parameter := ctx.Param("id")
	var id int
	id, err = strconv.Atoi(parameter)

	if err != nil {
		handler.ErrorBadReq(ctx)
		return
	}

	req.Id = id

	var resp model.WebResponse
	resp.Message = "Success Update Book"
	resp.Data, err = bc.BookService.Update(ctx, req)
	if err != nil {
		handler.ErrorNotFound(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (bc *BookControllerImpl) DeleteBook(ctx *gin.Context) {
	var response model.WebResponse
	parameter := ctx.Param("id")
	var id int
	id, err := strconv.Atoi(parameter)

	if err != nil {
		handler.ErrorBadReq(ctx)
		return
	}

	err = bc.BookService.Delete(ctx, id)
	if err != nil {
		handler.ErrorNotFound(ctx, err.Error())
		return
	}

	response.Message = "Sukses Menghapus Data Buku"
	response.Data = nil
	ctx.JSON(http.StatusOK, response)
}

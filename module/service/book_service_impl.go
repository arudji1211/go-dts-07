package service

import (
	"github.com/arudji1211/go-dts-07/module/handler"
	"github.com/arudji1211/go-dts-07/module/model"
	"github.com/arudji1211/go-dts-07/module/repository"
	"github.com/gin-gonic/gin"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(BookRepo repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: BookRepo,
	}
}

func (bs *BookServiceImpl) Insert(ctx *gin.Context, request model.BookCreateRequest) model.BookResponse {
	var data model.Book
	data.Author = request.Author
	data.Desc = request.Desc
	data.Title = request.Title
	data = bs.BookRepository.CreateBook(data)
	return handler.BookDomainToResp(data)
}

func (bs *BookServiceImpl) Update(ctx *gin.Context, request model.BookUpdateRequest) (model.BookResponse, error) {
	var data model.Book
	data.Author = request.Author
	data.Desc = request.Desc
	data.Title = request.Title
	data.Id = request.Id

	_, err := bs.BookRepository.GetBookById(data.Id)
	if err != nil {
		return handler.BookDomainToResp(data), err
	}

	data, err = bs.BookRepository.UpdateBook(data)
	if err != nil {
		return handler.BookDomainToResp(data), err
	}

	return handler.BookDomainToResp(data), nil
}

func (bs *BookServiceImpl) Delete(ctx *gin.Context, id int) error {
	err := bs.BookRepository.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookServiceImpl) GetAll(ctx *gin.Context) []model.BookResponse {
	data, err := bs.BookRepository.GetAllBook()
	if err != nil {
		panic(err)
	}

	var resp []model.BookResponse
	for _, vals := range data {
		resp = append(resp, handler.BookDomainToResp(vals))
	}
	return resp
}

func (bs *BookServiceImpl) GetById(ctx *gin.Context, id int) (model.BookResponse, error) {

	data, err := bs.BookRepository.GetBookById(id)

	if err != nil {
		return handler.BookDomainToResp(data), err
	}

	return handler.BookDomainToResp(data), err
}

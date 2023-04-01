package repository

import "github.com/arudji1211/go-dts-07/module/model"

type BookRepository interface {
	GetAllBook() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
	CreateBook(model.Book) (model.Book, error)
	UpdateBook(model.Book) error
	DeleteBook(id int) error
}

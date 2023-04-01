package repository

import (
	"errors"

	"github.com/arudji1211/go-dts-07/module/model"
)

type BookRepositoryImpl struct {
	Data map[uint64]model.Book
}

func NewBookRepository() BookRepository {
	DataStore := map[uint64]model.Book{
		0: {
			Id:      0,
			Author:  "aruji",
			Title:   "PHP",
			Desc:    "Buku PHP",
			Deleted: false,
		},
		1: {
			Id:      1,
			Author:  "hermatyar",
			Title:   "Python",
			Desc:    "Buku Python",
			Deleted: false,
		},
		2: {
			Id:      2,
			Author:  "gk tau siapa",
			Title:   "Java",
			Desc:    "Buku Java",
			Deleted: false,
		},
	}
	return &BookRepositoryImpl{
		Data: DataStore,
	}
}

// b BookRepositoryImpl
func (b *BookRepositoryImpl) GetAllBook() ([]model.Book, error) {
	var data []model.Book
	for _, vals := range b.Data {
		data = append(data, vals)
	}
	return data, nil
}

func (b *BookRepositoryImpl) GetBookById(id int) (model.Book, error) {
	var data model.Book
	valid := false

	for _, vals := range b.Data {
		if id == int(vals.Id) && vals.Deleted != true {
			data = model.Book{
				Id:     vals.Id,
				Title:  vals.Title,
				Author: vals.Author,
				Desc:   vals.Desc,
			}
			valid = true
			break
		}
	}

	if valid {
		return data, nil
	} else {
		err := errors.New("Data tidak di temukan")
		return data, err
	}
}

func (b *BookRepositoryImpl) CreateBook(book model.Book) model.Book {
	counter := len(b.Data)
	book.Id = counter
	book.Deleted = false
	b.Data[uint64(counter)] = book
	return b.Data[uint64(counter)]
}

func (b *BookRepositoryImpl) UpdateBook(book model.Book) (model.Book, error) {
	if b.Data[uint64(book.Id)].Deleted {
		err := errors.New("Data tidak di temukan")
		return book, err
	} else {
		b.Data[uint64(book.Id)] = book
		return book, nil
	}
}

func (b *BookRepositoryImpl) DeleteBook(id int) error {

	if id < len(b.Data)-1 && b.Data[uint64(id)].Deleted == false {
		book := b.Data[uint64(id)]
		book.Deleted = true
		b.Data[uint64(id)] = book
		return nil
	} else {
		return errors.New("Data tidak di temukan")
	}
}

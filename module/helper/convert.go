package helper

import "github.com/arudji1211/go-dts-07/module/model"

func BookDomainToResp(data model.Book) model.BookResponse {
	resp := model.BookResponse{
		Id:     data.Id,
		Title:  data.Title,
		Author: data.Author,
		Desc:   data.Desc,
	}
	return resp
}

package model

type BookCreateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

package quiz

type MapMurid map[string]string

type Murid interface {
	GetMurid(id int) MapMurid
	ShowMurid(data MapMurid)
}

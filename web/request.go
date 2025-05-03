package web

type Request struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

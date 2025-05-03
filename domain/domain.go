package domain

type Domain struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

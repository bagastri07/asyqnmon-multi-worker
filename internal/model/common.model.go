package model

type PageData struct {
	Title string
	Items []MenuItem
}

type MenuItem struct {
	Name string
	URL  string
}

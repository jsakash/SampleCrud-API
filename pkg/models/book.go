package models

type Book struct {
	Id     int    `json:"id" gorm:"primarKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

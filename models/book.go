package models

type Book struct {
	ISBN   uint   `json:"isbn" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
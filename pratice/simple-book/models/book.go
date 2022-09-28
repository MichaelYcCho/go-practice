package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key; auto_increment"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type User struct {
	ID uint `json:"id" gorm:"primary_key; auto_increment"`
}

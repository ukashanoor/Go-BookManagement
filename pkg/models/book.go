package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ukashanoor/Go-BookManagement/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id uint64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID =?", id).First(&book)
	return &book, db
}

func DeleteBook(id uint64) Book {
	var book Book
	db.Where("ID =?", id).Delete(book)
	return book
}

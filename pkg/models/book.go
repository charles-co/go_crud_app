package models

import (
	"github.com/charles-co/go_crud_app/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(&b).Error; err != nil {
		return &Book{}, err
	}
	return b, nil
}

func (b *Book) GetBooks() (*[]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return &[]Book{}, err
	}
	return &books, nil
}

func (b *Book) GetBook(id string) (*Book, error) {
	var book Book
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return &Book{}, err
	}
	return &book, nil
}

func (b *Book) UpdateBook(id string) (*Book, error) {
	db.Save(&b)
	return b, nil
}

func (b *Book) DeleteBook(id string) (int64, error) {
	db.Where("id = ?", id).Delete(&b)
	return db.RowsAffected, nil
}

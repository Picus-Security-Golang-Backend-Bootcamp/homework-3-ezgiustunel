package book

import (
	"errors"
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)

	return books
}

func (b *BookRepository) Migration() {
	b.db.AutoMigrate(&Book{})
}

func (b *BookRepository) InsertData(books []models.Book) {
	for _, book := range books {
		b.db.Where(Book{ID: book.Id}).Attrs(Book{StockNumber: book.StockNumber, PageNumber: book.PageNumber, Price: book.Price, Name: book.Name, StockCode: book.StockCode, Isbn: book.Isbn, AuthorName: book.AuthorName}).FirstOrCreate(&book)
	}
}

func (b *BookRepository) FindByBookName(bookName string) []Book {
	var cities []Book
	b.db.Where("Name LIKE ?", "%"+bookName+"%").Find(&cities)
	return cities
}

func (b *BookRepository) GetById(id int) Book {
	var book Book
	result := b.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Book not found with id : %d", id)
		return Book{}
	}
	return book
}

func (b *BookRepository) Update(book Book) error {
	result := b.db.Save(book)
	//b.db.Model(&book).Update("name", "deneme")

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) DeleteById(id int) error {
	result := b.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

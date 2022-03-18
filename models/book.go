package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id          int
	StockNumber int
	PageNumber  int
	Price       float64
	Name        string
	StockCode   string
	Isbn        string
	AuthorName  string
}

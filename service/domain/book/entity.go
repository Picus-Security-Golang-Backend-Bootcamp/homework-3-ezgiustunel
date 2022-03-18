package book

type Book struct {
	ID          int `gorm:"primaryKey;autoIncrement" json:"Id"` //Change it
	StockNumber int
	PageNumber  int
	Price       float64
	Name        string
	StockCode   string
	Isbn        string
	AuthorName  string
}

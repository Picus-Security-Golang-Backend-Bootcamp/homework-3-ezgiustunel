package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/service/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/service/infrastructure"
)

// var books, authors []string

// func init() {
// 	//book list
// 	books = []string{"Simyaci",
// 		"Bab-i Esrar",
// 		"Nar Ağaci",
// 		"Fareler ve İnsanlar",
// 		"Kürk Mantolu Madonna",
// 		"Hayvan Çiftliği",
// 		"Şeker Portakali",
// 		"Uçurtma Avcisi",
// 		"Suç ve Ceza",
// 		"Serenad",
// 		"Yeraltindan Notlar",
// 		"Toprak Ana",
// 		"Fatih Harbiye",
// 		"Saatleri Ayarlama Enstitüsü",
// 		"Acimak",
// 		"Ateşten Gömlek",
// 		"Çocukluğum",
// 		"Aşk",
// 		"Kuyucakli Yusuf",
// 		"Arkadaş",
// 		"Momo",
// 	}

// 	//author list
// 	authors = []string{"Paulo Coelho",
// 		"Ahmet Ümit",
// 		"Nazan Bekiroğlu",
// 		"John Steinback",
// 		"Sabahattin Ali",
// 		"George Orwell",
// 		"Mauro Vasgencelos",
// 		"Halid Hüseyni",
// 		"Fyodor Dostoyevski",
// 		"Zülfü Livaneli",
// 		"Fyodor Dostoyevski",
// 		"Cengiz Aytmatov",
// 		"Peyami Safa",
// 		"Ahmet Hamdi Tanpinar",
// 		"Reşat Nuri Güntekin",
// 		"Halide Edip Adivar",
// 		"Maksim Gorki",
// 		"Elif Şafak",
// 		"Sabahattin Ali",
// 		"Gorki",
// 		"Michael Ende",
// 	}
// }
var bookList []models.Book
var repository *book.BookRepository

const list = "list"
const search = "search"
const buy = "buy"
const delete = "delete"

func init() {
	db := infrastructure.ConnectDB("postgres://ezgiustunel:pass@localhost:5432/library")
	repository = book.NewBookRepository(db)
	repository.Migration()

	bookList, _ = helper.ReadCsv("book.csv")
	repository.InsertData(bookList)
}

func main() {
	args := os.Args

	if len(args) == 1 {
		helper.PrintMessagesToConsole()
		return
	}

	firstInput := args[1]

	switch firstInput {
	case list:
		listBooks()
	case search:
		searchBook(args)
	case buy:
		buyBook(args)
	case delete:
		deleteBook(args)
	default:
		helper.PrintMessagesToConsole()
	}
}

func listBooks() {
	books := repository.FindAll()

	for _, book := range books {
		fmt.Println(book.Name)
	}
}

func searchBook(args []string) {
	if len(args) < 3 {
		helper.PrintMessagesToConsole()
		return
	}

	searchedBook := strings.Join(args[2:], " ")
	books := repository.FindByBookName(searchedBook)

	for _, book := range books {
		fmt.Printf("name: %s, author: %s\n", book.Name, book.AuthorName)
	}
}

func buyBook(args []string) {
	if len(args) < 4 {
		helper.PrintMessagesToConsole()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Please give a number")
		return
	}

	bookNumber, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Please give a number")
		return
	}

	book := repository.GetById(id)

	book.StockNumber -= bookNumber

	repository.Update(book)

}

func deleteBook(args []string) {
	if len(args) != 3 {
		helper.PrintMessagesToConsole()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Please give a number")
		return
	}

	repository.DeleteById(id)
}

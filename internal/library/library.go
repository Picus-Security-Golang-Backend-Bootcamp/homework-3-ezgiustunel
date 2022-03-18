package library

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/models"
)

func PrintList(bookList []models.Book) {
	if len(bookList) == 0 {
		fmt.Println("okuyamadım")
		return
	}
	fmt.Println()
	for _, value := range bookList {
		fmt.Println(value.Name)
	}
	fmt.Println()
}

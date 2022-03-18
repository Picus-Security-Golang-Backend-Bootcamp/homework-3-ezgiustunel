package helper

import (
	"encoding/csv"
	"os"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/models"
)

func ReadCsv(filename string) ([]models.Book, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []models.Book

	for _, line := range lines[1:] {
		id, _ := ConvertStringToInt(line[0])
		stockNumber, _ := ConvertStringToInt(line[2])
		pageNumber, _ := ConvertStringToInt(line[3])
		price, _ := ConvertStringToFloat64(line[4])

		data := models.Book{
			Id:          id,
			Name:        line[1],
			StockNumber: stockNumber,
			PageNumber:  pageNumber,
			Price:       price,
			StockCode:   line[5],
			Isbn:        line[6],
			AuthorName:  line[7],
		}
		result = append(result, data)
	}

	return result, nil
}

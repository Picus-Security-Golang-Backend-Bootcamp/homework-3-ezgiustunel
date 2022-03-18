package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

// Method for converting string input to int
// if it is a number then return the number, otherwise return -1
func ConvertStringToInt(input string) (int, error) {
	result, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func ConvertStringToFloat64(input string) (float64, error) {
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func GenerateRandomInt(max int64) (int, error) {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(max))
	bigStr := fmt.Sprint(bigInt)
	result, err := ConvertStringToInt(bigStr)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func GenerateRandomFloat(max int64) float64 {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(max))
	bigStr := fmt.Sprint(bigInt)
	result, _ := ConvertStringToFloat64(bigStr)
	if result != -1 {
		return result
	}
	fmt.Printf("\nUnexpected error")
	return 0
}

func GenerateRandomCode(length int) (string, error) {
	seed := "012345679"
	byteSlice := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}

// Method for print messages to user in the console.
func PrintMessagesToConsole() {
	fmt.Printf("\n--Invalid Input--\n\n")
	fmt.Println("You can use the methods below to make some actions on book list")
	fmt.Println("list: Lists the books")
	fmt.Println("search \"bookname\": searches the bookname given in the book list")
	fmt.Println("buy: you can buy books")
	fmt.Printf("delete: you can delete a book from book list\n\n")
}

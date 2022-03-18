package infrastructure

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(conString string) *gorm.DB {
	dbURL := "postgres://ezgiustunel:pass@localhost:5432/library"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt: true, // sonraki sorgular için cache
	})

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database : %s", err.Error()))
	}

	db.AutoMigrate(&models.Book{})
	fmt.Println(db)

	if err != nil {
		panic(err)
	}

	return db
}

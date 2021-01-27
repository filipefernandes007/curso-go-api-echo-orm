package main

import (
	"curso-go/api-echo-orm/database"
	"curso-go/api-echo-orm/handlers"
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

const HttpApiPort = "9005"

func main() {
	db, err := database.ConnectAndMigrate()
	if err != nil {
		panic(err)
	}

	CreateRoutes(db)
}

func CreateRoutes(db *gorm.DB) {
	e := echo.New()
	internalHandler := handlers.NewInternalHandler(e)
	internalHandler.Healthy()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", HttpApiPort)))
}

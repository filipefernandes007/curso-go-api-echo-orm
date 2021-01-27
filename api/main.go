package main

import (
	"curso-go/api-echo-orm/handlers"
	"fmt"
	"github.com/labstack/echo"
)

const HttpApiPort = "9005"

func main() {
	CreateRoutes()
}

func CreateRoutes() {
	e := echo.New()
	internalHandler := handlers.NewInternalHandler(e)
	internalHandler.Healthy()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", HttpApiPort)))
}

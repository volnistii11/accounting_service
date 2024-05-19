package main

import (
	"log"

	"github.com/volnistii11/accounting_service/transfer/internal/app"
)

func main() {
	application := app.NewApp()

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}

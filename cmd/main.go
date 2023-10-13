package main

import (
	"auction-be/internal/app"
	"os"
)

func main() {
	application := app.NewApp()
	if err := application.StartServer(":8080"); err != nil {
		os.Exit(0)
	}
}

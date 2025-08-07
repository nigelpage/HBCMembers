package main

import (
	"log"

	"github.com/nigelpage/HBCMembers/internal/application"
)

func main() {
	app := application.NewApp()
	if err := app.Run(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}

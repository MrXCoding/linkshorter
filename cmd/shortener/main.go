package main

import (
	"github.com/MrXCoding/linkshorter/internal/storage"
	"log"

	"github.com/MrXCoding/linkshorter/cmd/server"
)

func main() {
	db := storage.NewInMemory()

	err := server.Run(db)
	if err != nil {
		log.Fatal(err)
	}
}

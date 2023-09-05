package main

import (
	"github.com/MrXCoding/linkshorter/internal/storage"
	"github.com/MrXCoding/linkshorter/pkg/hasher"
	"log"

	"github.com/MrXCoding/linkshorter/cmd/server"
)

func main() {
	db := storage.NewInMemory(&hasher.Sha256Base68{})

	err := server.Run(db)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/MrXCoding/linkshorter/internal/config"
	"github.com/MrXCoding/linkshorter/internal/storage"
	"github.com/MrXCoding/linkshorter/pkg/hasher"
	"log"

	"github.com/MrXCoding/linkshorter/cmd/server"
)

func init() {
	config.Parse()
}

func main() {
	db := storage.NewMap(&hasher.Sha256Base68{})
	conf := config.New()

	err := server.Run(db, conf)
	if err != nil {
		log.Fatal(err)
	}
}

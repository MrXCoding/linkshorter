package server

import (
	"net/http"

	"github.com/MrXCoding/linkshorter/internal/handlers"
	"github.com/MrXCoding/linkshorter/internal/storage"
)

const serverAddr = `:8080`

func Run(db storage.Repository) error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handlers.Main(db))

	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		return err
	}

	return nil
}

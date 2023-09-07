package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/MrXCoding/linkshorter/internal/handlers"
	"github.com/MrXCoding/linkshorter/internal/storage"
)

const serverAddr = `:8080`

func Run(db storage.Repository) error {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/{hash}", handlers.Get(db))
		r.Post("/", handlers.Save(db))
	})

	err := http.ListenAndServe(serverAddr, r)
	if err != nil {
		return err
	}

	return nil
}

package server

import (
	"github.com/MrXCoding/linkshorter/internal/config"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/MrXCoding/linkshorter/internal/handlers"
	"github.com/MrXCoding/linkshorter/internal/storage"
)

func Run(db storage.Repository, config config.Main) error {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", handlers.Save(db, config))
		r.Get("/{hash}", handlers.Get(db))
	})

	err := http.ListenAndServe(config.Host(), r)
	if err != nil {
		return err
	}

	return nil
}

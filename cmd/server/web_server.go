package server

import (
	"net/http"

	"github.com/MrXCoding/linkshorter/internal/handlers"
	"github.com/MrXCoding/linkshorter/internal/storage"
	"github.com/MrXCoding/linkshorter/internal/validation"
)

const serverAddr = `:8080`

func Run(db storage.Repository) error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handle(db))

	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		return err
	}

	return nil
}

func handle(db storage.Repository) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if isValid, err := validation.Validate(req); !isValid {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		switch req.Method {
		case http.MethodPost:
			handlers.Save(db)(res, req)
		case http.MethodGet:
			handlers.Get(db)(res, req)
		}
	}
}

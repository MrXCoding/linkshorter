package handlers

import (
	"io"
	"net/http"

	"github.com/MrXCoding/linkshorter/internal/storage"
)

const baseURL = "http://localhost:8080/"

func Save(db storage.Repository) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		url, err := io.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusCreated)
			res.Write([]byte(""))
			return
		}

		hash, err := db.Save(string(url))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("unkonwn url"))
		}

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(baseURL + hash))
	}
}

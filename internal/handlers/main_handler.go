package handlers

import (
	"github.com/MrXCoding/linkshorter/internal/storage"
	"github.com/MrXCoding/linkshorter/internal/validation"
	"io"
	"net/http"
	"strings"
)

const baseURL = "http://localhost:8080/"

func Main(db storage.Repository) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if isValid, err := validation.Validate(req); !isValid {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		switch req.Method {
		case http.MethodPost:
			save(db)(res, req)
		case http.MethodGet:
			get(db)(res, req)
		}
	}
}

func get(db storage.Repository) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		arr := strings.Split(req.URL.Path, "/")
		hash := arr[1]

		url, err := db.Get(hash)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
		}

		res.Header().Add("Location", url)
		res.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func save(db storage.Repository) http.HandlerFunc {
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

package handlers

import (
	"github.com/MrXCoding/linkshorter/internal/config"
	"github.com/MrXCoding/linkshorter/internal/storage"
	"io"
	"net/http"
	"strings"
)

func Get(db storage.Repository) http.HandlerFunc {
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

func Save(db storage.Repository, config config.Main) http.HandlerFunc {
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
		res.Write([]byte(config.GetBaseURL() + hash))
	}
}

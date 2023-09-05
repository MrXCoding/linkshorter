package handlers

import (
	"net/http"

	"github.com/MrXCoding/linkshorter/internal/storage"
)

func Get(db storage.Repository) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		hash := req.URL.Path[1:]

		url, err := db.Get(hash)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("unkonwn hash"))
		}

		res.Header().Add("Location", url)
		res.WriteHeader(http.StatusTemporaryRedirect)
	}
}

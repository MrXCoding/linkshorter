package handlers

import (
	"github.com/MrXCoding/linkshorter/internal/storage"
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

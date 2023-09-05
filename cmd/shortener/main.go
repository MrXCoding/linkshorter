package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

const (
	baseURL       = "http://localhost:8080/"
	deafaultHashe = "EwHXdJfB"
	contentType   = "text/plain"
)

var (
	methodAllowedError = "Only POST and GET methods allowed"
	contentTypeError   = "Only text/plain header allowed"
)

var hashStorage = map[string]string{}

func handle(res http.ResponseWriter, req *http.Request) {
	if isValid, err := validate(req); !isValid {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	process(res, req)
}

func validate(req *http.Request) (bool, error) {
	if req.Method != http.MethodPost && req.Method != http.MethodGet {
		return false, errors.New(methodAllowedError)
	}

	if ctype := req.Header.Get("Content-Type"); ctype != contentType {
		return false, errors.New(contentTypeError)
	}

	return true, nil
}

func process(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		save(res, req)
	case http.MethodGet:
		get(res, req)
	}
}

func save(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	url, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Error body reading"))
	}

	hashStorage[deafaultHashe] = string(url)

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(baseUrl + deafaultHashe))
}

func get(res http.ResponseWriter, req *http.Request) {
	hash := req.URL.Path[1:]

	url, ok := hashStorage[hash]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("unkonwn hash"))
	}

	res.Header().Add("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handle)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

const (
	BASE_URL      = "http://localhost:8080/"
	DEFAULT_HASHE = "EwHXdJfB"
	CONTENT_TYPE  = "text/plain"
)

var (
	MethodAllowedError = "Only POST and GET methods allowed"
	ContentTypeError   = "Only text/plain header allowed"
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
		return false, errors.New(MethodAllowedError)
	}

	if ctype := req.Header.Get("Content-Type"); ctype != CONTENT_TYPE {
		return false, errors.New(ContentTypeError)
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

	hashStorage[DEFAULT_HASHE] = string(url)

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(BASE_URL + DEFAULT_HASHE))
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

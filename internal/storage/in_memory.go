package storage

import (
	"errors"
	"github.com/MrXCoding/linkshorter/pkg/hasher"
)

type InMemory struct {
	storage map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{
		storage: make(map[string]string),
	}
}

func (im *InMemory) Save(url string) (string, error) {
	hash := hasher.Encode(url, "")

	im.storage[hash] = url

	return hash, nil
}

func (im *InMemory) Get(hash string) (string, error) {
	url, ok := im.storage[hash]
	if !ok {
		return "", errors.New(hash + " not found")
	}

	return url, nil
}

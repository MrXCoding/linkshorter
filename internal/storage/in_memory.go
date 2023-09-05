package storage

import (
	"errors"
	"github.com/MrXCoding/linkshorter/pkg/hasher"
)

type InMemory struct {
	hasher  hasher.Generator
	storage map[string]string
}

func NewInMemory(hasher hasher.Generator) *InMemory {
	return &InMemory{
		hasher:  hasher,
		storage: make(map[string]string),
	}
}

func (im *InMemory) Save(url string) (string, error) {
	hash := im.hasher.Encode(url, "")

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

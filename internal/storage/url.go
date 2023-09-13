package storage

import (
	"errors"
	"github.com/MrXCoding/linkshorter/pkg/hasher"
)

type Map struct {
	hasher  hasher.Generator
	storage map[string]string
}

func NewMap(hasher hasher.Generator) *Map {
	return &Map{
		hasher:  hasher,
		storage: make(map[string]string),
	}
}

func (im *Map) Save(url string) (string, error) {
	hash := im.hasher.Encode(url, "")

	im.storage[hash] = url

	return hash, nil
}

func (im *Map) Get(hash string) (string, error) {
	url, ok := im.storage[hash]
	if !ok {
		return "", errors.New(hash + " not found")
	}

	return url, nil
}

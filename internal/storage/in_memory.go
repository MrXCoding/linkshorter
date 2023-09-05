package storage

type InMemory struct {
	storage map[string]string
}

const hash = "EwHXdJfB"

func NewInMemory() *InMemory {
	return &InMemory{
		storage: make(map[string]string),
	}
}

func (im *InMemory) Save(url string) (string, error) {
	im.storage[hash] = url

	return hash, nil
}

func (im *InMemory) Get(hash string) (string, error) {
	url, ok := im.storage[hash]
	if !ok {
		return "", nil
	}

	return url, nil
}

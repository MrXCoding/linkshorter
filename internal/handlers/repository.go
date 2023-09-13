package handlers

type Repository interface {
	Save(url string) (string, error)
	Get(hash string) (string, error)
}

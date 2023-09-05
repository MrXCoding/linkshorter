package hasher

type Generator interface {
	Encode(str string, seed string) string
}

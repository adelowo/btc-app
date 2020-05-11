package queryapp

type Client interface {
	FetchPrice() (float64, error)
}

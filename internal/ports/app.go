package ports

type APIPort interface {
	GetShortenUrl(url string, userId string) (string, error)
}

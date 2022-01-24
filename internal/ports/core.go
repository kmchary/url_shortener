package ports

// the core functionality of this application is to provide shortened url
type UrlShortenerPort interface {
	ShortenUrl(url string, userId string) (string, error)
}

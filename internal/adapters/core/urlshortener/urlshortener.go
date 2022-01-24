package urlshortener

type Adapter struct {
}

func (ad Adapter) ShortenUrl(url string, userId string) (string, error) {
	return url + userId, nil
}

func NewAdapter() *Adapter {
	return &Adapter{}
}
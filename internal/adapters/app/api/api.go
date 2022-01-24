package api

import (
	"github.com/kmchary/url_shortener/internal/ports"
)

type Adapter struct {
	us ports.UrlShortenerPort
}

func (apia Adapter) ShortenUrl(url string, userId string) (string, error)  {
	return  apia.ShortenUrl(url, userId)
}

func NewAdapter(us ports.UrlShortenerPort) *Adapter {
	return &Adapter{ us: us	}
}

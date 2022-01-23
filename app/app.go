package app

import (
	"net/http"

	"github.com/kmchary/url_shortener/router"
)

type UrlShortenerApplicationInterface interface {
	Init() error
	Start(serverPort string) error
}

type UrlShortenerApplication struct {
	router             router.RouterInterface
	server             *http.Server
	diServiceContainer *di.ServiceContainer
}

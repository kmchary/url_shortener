package dicontainer

import (
	"context"

	caching "github.com/kmchary/url_shortener/cachingstore"
	"github.com/kmchary/url_shortener/controller"
)

type DependencyInjectionContainer struct {
	Controller   *controller.UrlShortenerController
	CachingStore caching.CachingStore
}

func (c *DependencyInjectionContainer) StartDependencyInjection() {
	c.initRedisClient()
}

func (c *DependencyInjectionContainer) initRedisClient() {
	c.CachingStore = caching.NewCashingStore(context.Background(), "localhost", "6379")
}

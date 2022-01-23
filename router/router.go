package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kmchary/url_shortener/controller"
)

type RouterInterface interface {
	InitRoutes(controller.UrlShortenerController)
	GetMux() *gin.Engine
}

type router struct {
	mux *gin.Engine
}

func NewRouter() RouterInterface {
	return &router{mux: gin.Default()}
}

func (r *router) GetMux() *gin.Engine {
	return r.mux
}

func (r *router) InitRoutes(c controller.UrlShortenerController) {
	r.mux.POST("/shorten", c.ShortenURL)
}

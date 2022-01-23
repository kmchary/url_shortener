package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmchary/url_shortener/service"
)

type UrlShortenerController struct {
	urlSHortenerService service.UrlShortenerServiceInterface
	cachingService      service.CacheServiceInterface
}

type Request struct {
	Url    string `form:"url" json:"url" binding:"required"`
	UserId string `form:"user_id" json:"user_id" binding:"required"`
}

func (uc *UrlShortenerController) ShortenURL(c *gin.Context) {
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	surl := uc.urlSHortenerService.ShortenURL(request.Url, request.UserId)

	_ = uc.cachingService.SetNX(ctx, string(encodedString), url)
	c.JSON(http.StatusOK, gin.H{"original_url": request.Url, "shortened_url": "test.com" + "/" + surl})
}

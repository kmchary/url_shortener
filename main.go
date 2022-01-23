package main

import (
	"github.com/kmchary/url_shortener/router"
)

// const (
// 	domainName = "www.testurlshortner.com"
// )

//var ctx = context.Background()

func main() {
	r := router.NewRouter()
	r.InitRoutes()
	r.GetMux().Run("8080")
}

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rusisg/atyrau-news-go/controller"
)

func main() {
	r := gin.Default()

	r.Static("/css", "public/assets/css")
	r.Static("/logo", "public/assets/img")
	r.LoadHTMLGlob("public/templates/*.html")

	r.GET("/", controller.IndexGet)
	r.GET("/news", controller.NewsGet)
	r.GET("/news_detail", controller.NewsDetailGet)

	// scraper...

	if err := r.Run(":3000"); err != nil {
		fmt.Printf("Error while run the server %v", err)
	}
}

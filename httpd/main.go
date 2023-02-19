package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rusisg/atyrau-news-go/httpd/handler"
)

func main() {
	r := gin.Default()

	r.Static("/css", "./public/css")
	r.Static("/logo", "./public/logo")
	r.LoadHTMLGlob("./public/html/*.html")

	r.GET("/", handler.IndexGet)
	r.GET("/news", handler.NewsGet)
	r.GET("/news_detail", handler.NewsDetailGet)

	// scraper...

	if err := r.Run(":3000"); err != nil {
		fmt.Printf("Error while run the server %v", err)
	}
}

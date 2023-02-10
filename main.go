package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	// tmpl := template.Must(template.ParseFiles("public/html/base.html", "public/html/header.html", "public/html/index.html", "public/html/news.html", "public/html/news_detail.html"))

	r.Static("/css", "./public/css")
	r.Static("/logo", "./public/logo")
	r.LoadHTMLGlob("./public/html/*.html")

	r.GET("/news", news)
	r.GET("/news_detail", news_detail)
	r.GET("/", index)

	if err := r.Run(":3000"); err != nil {
		fmt.Printf("Error while run the server %v", err)
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title": "Atyrau Bugin",
	})

}

func news(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title": "Latest News",
		"articles": []interface{}{
			map[string]interface{}{
				"title":   "Article 1",
				"content": "This is the first article",
			},
			map[string]interface{}{
				"title":   "Article 2",
				"content": "This is the second article",
			},
		},
	})
}

func news_detail(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title": "Atyrau Bugin",
	})
}

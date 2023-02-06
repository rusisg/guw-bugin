package main

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	tmpl := template.Must(template.ParseFiles("public/html/base.html", "public/html/header.html", "public/html/index.html", "public/html/news.html", "public/html/news_detail.html"))

	r.Static("/css", "./public/css")
	r.LoadHTMLGlob("./public/html/*.html")

	r.GET("/news", news)
	r.GET("/", func(c *gin.Context) {
		renderTemplate(c, tmpl, gin.H{
			"title": "Atyrau Bugin",
		})
	})

	r.Run(":3000")
}

func news(c *gin.Context) {
	c.HTML(http.StatusOK, "news.tmpl", gin.H{
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

func renderTemplate(c *gin.Context, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(c.Writer, data)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/css", "./public/css")
	r.LoadHTMLFiles("./public/html/index.html")

	r.GET("/", index)
	r.Run(":3000")
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

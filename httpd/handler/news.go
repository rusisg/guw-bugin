package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsGet(c *gin.Context) {
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

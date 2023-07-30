package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsDetailGet(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title": "Atyrau Bugin",
	})
}

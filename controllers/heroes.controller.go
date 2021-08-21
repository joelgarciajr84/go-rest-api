package controllers

import "github.com/gin-gonic/gin"

func ShowBestHero(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "Batman",
	})
}

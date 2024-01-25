package main

import (
	"github.com/gin-gonic/gin"
)

func GetContainer(c *gin.Context) {
	c.HTML(200, "container.html", gin.H {})
}

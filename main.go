package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accetpt")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	r.Run(":8000")
}

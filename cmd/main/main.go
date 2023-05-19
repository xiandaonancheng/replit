package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/description", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"ok",
			"data":map[string]interface{}{
				"description":"this is replit api service",
			},
		})
	})
	server.Run(":80")
}
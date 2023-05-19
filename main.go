package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	server := gin.Default()
	server.GET("/description", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":0,
			"msg":"ok",
			"data":map[string]interface{}{
				"description":"this is replit api service",
				"time":time.Now().Format("2006-01-02 15:04:05.000000000 -0700 "),
			},
		})
	})
	server.Run(":80")
}
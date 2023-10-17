package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello,gin",
		})
	})
	r.GET("/page1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page1.html", gin.H{
			"title": "hello,page1",
		})
	})
	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}

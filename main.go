package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("template/demo.html") // html 제공
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo.html", gin.H{
			"title": "IDIS Demo Page",
		})
	})

	// 서버 실행
	router.Run(":8090") // 8080 포트에서 서버 실행
}

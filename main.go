package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./template2")     // js 등 파일 제공
	router.LoadHTMLFiles("template2/demo.html") // html 제공
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo.html", gin.H{
			"title": "Preference's Test Page",
		})
	})

	router.POST("/submit-result", func(c *gin.Context) {
		var result ExamResult
		if err := c.ShouldBindJSON(&result); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 이메일 전송
		if err := sendEmail(result); err != nil {
			log.Printf("Failed to send email: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Result sent successfully"})
	})

	// certFile := "public.pem"
	certFile := "/etc/letsencrypt/live/jy-onestaste-29.n-e.kr/fullchain.pem"
	// keyFile := "private.pem"
	keyFile := "/etc/letsencrypt/live/jy-onestaste-29.n-e.kr/privkey.pem"

	if err := router.RunTLS(":443", certFile, keyFile); err != nil {
		log.Fatal("Unable to start Server : ", err)
	}

	// 서버 실행
	// router.Run(":8090") // 8080 포트에서 서버 실행
}

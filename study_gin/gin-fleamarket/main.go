package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // gin.Default()는 기본적으로 Logger와 Recovery 미들웨어를 포함한 Gin 엔진을 생성한다. HTTP 요청을 받고 응답을 전송하는 역할을 한다.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}) // GET 요청을 받고, "/ping" 경로로 요청이 들어오면, JSON 형식으로 응답을 보내준다.
	r.Run("localhost:8080") // 8080 포트로 서버를 실행한다.
}

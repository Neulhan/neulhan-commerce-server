package middleware

import (
	"github.com/gin-gonic/gin"
)

func CustomMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 전 실행
		//fmt.Println("----------------")
		// 요청 실행
		c.Next()
		// 요청 후 실행
		//fmt.Println("----------------")
	}
}

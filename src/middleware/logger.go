package middleware

import (
	"github.com/kataras/iris/v12"
)

func Logger() iris.Handler {
	return func(c iris.Context) {
		// 요청 전 실행
		//fmt.Println("----------------")
		// 요청 실행
		c.Next()
		// 요청 후 실행
		//fmt.Println("----------------")
	}
}

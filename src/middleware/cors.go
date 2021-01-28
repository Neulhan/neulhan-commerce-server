package middleware

import "github.com/kataras/iris/v12"

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if ctx.Method() == iris.MethodOptions {
		ctx.Header("Access-Control-Methods", "POST, PUT, PATCH, DELETE")
		ctx.Header("Access-Control-Allow-Methods", "POST, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")
		ctx.Header("Access-Control-Max-Age", "86400")
		ctx.StatusCode(iris.StatusNoContent)
		return
	}

	ctx.Next()
}

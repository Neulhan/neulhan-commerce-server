package middleware

import (
	"github.com/kataras/iris/v12"
	"log"
	"neulhan-commerce-server/src/jwt"
)

func UserMiddleware() iris.Handler {
	return func(c iris.Context) {
		accessToken := c.GetCookie("accessToken")
		log.Println(accessToken)
		if accessToken == "" {
			c.Next()
			return
		}
		var err error
		claims := &jwt.Claims{}
		claims, err = jwt.ParseToken(accessToken)
		if err != nil {
			c.StopWithError(iris.StatusForbidden, err)
		}
		c.Values().Set("UserID", claims.UserID)
		c.Next()
	}
}

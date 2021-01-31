package auth

import (
	"github.com/kataras/iris/v12"
)

type KakaoData struct {
	ID        int    `json:"kakaoID"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	FullImage string `json:"fullImage"`
}

func KakaoAuth(ctx iris.Context) (err error) {
	var kakaoData KakaoData
	err = ctx.ReadJSON(&kakaoData)
	if err != nil {
		return err
	}
	ctx.SetCookieKV("k", "v")
	return nil
}

func GithubAuth(id string) {

}

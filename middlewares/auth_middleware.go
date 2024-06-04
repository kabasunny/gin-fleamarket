package middlewares

import (
	"gin-fleamarket/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization") //POST時に、Bearer Tokenの項目に、ログインで取得したトークン文字列を入力すると、ここで抽出する模様
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized) //後続の処理の中止
			return //現在の処理の終了
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return // POST時にBearer Tokenとなっていないといけない
		}

		tokenString := strings.TrimPrefix(header, "Bearer ") 
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user) //リクエストの生存期間中に、"user"キーでユーザー情報を取得できる。

		ctx.Next()
	}
}
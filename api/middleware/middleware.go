package middleware

import (
	"apii_gateway/api/handlers/v1/tokens"
	"net/http"
	"apii_gateway/config"

	"github.com/gin-gonic/gin"
)


func Auth(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/v1/user/login" || ctx.Request.URL.Path == "/v1/user/register" {
		ctx.Next()
		return
	}

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	cfg := config.Load()

	_, err := tokens.ExtractClaim(token, []byte(cfg.SignInKey))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "authorized"},)
	ctx.Next()
}



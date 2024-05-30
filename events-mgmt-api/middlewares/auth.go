package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mobamoh/go-playground/events-mgmt-api/utils"
)

func Authenticate(ctx *gin.Context) {

	authzToken := ctx.Request.Header.Get("Authorization")
	if authzToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthirized operation"})
		return
	}

	uid, err := utils.VerifyToken(authzToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthirized operation"})
		return
	}
	ctx.Set("uid", uid)
	ctx.Next()
}

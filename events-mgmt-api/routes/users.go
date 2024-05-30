package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mobamoh/go-playground/events-mgmt-api/internal/models/user"
	"github.com/mobamoh/go-playground/events-mgmt-api/utils"
)

func Signup(ctx *gin.Context) {
	var user user.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid json format"})
		return
	}
	if err := user.Save(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "couldn't created user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "user created successfully"})
}

func Login(ctx *gin.Context) {
	var user user.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid json format"})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

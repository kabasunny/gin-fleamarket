package controllers

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := c.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.Status(http.StatusCreated)
}
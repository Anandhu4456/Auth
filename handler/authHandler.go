package handler

import (
	"auth/dto"
	"auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	CreateUser(ctx *gin.Context)
}

type AuthHandlerImpl struct {
	authService service.AuthService
}

// auth handler constructor
func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &AuthHandlerImpl{
		authService: authService,
	}
}

func (h *AuthHandlerImpl) CreateUser(ctx *gin.Context) {
	var userReq dto.UserCreateRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.authService.CreateUser(ctx.Request.Context(), &userReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

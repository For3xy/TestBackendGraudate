package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testbackendGraudate/domain"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input domain.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

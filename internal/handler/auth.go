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

func (h *Handler) signIn(ctx *gin.Context) {

}

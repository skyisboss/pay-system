package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"err": 0, "msg": "ok", "data": gin.H{}})
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ApiMiddle(ctx *gin.Context) {
	var req struct {
		AppID string `json:"appid" binding:"required" validate:"alphanum"`
		Nonce string `json:"nonce" binding:"required" validate:"min=6,max=42,alphanum"`
		Sign  string `json:"sign" binding:"required" validate:"len=32,alphanum"`
	}
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "参数不完整", "data": gin.H{}})
		ctx.Abort()
		return
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": err.Error(), "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 获取产品信息
}

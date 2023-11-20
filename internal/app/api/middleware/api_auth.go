package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/util"
)

// 签名验证
// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=20_1
func ApiMiddle(ioc *ioc.Container) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := context.Background()
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
		productRow, err := ioc.ProductService().GetByAppID(context, req.AppID)
		if ent.IsNotFound(err) || productRow == nil {
			ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "商户不存在", "data": gin.H{}})
			ctx.Abort()
			return
		}
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "内部错误", "data": gin.H{}})
			ctx.Abort()
			return
		}

		// 验证签名
		checkSign := util.WechatSign(productRow.AppSecret, gin.H{
			"appid": req.AppID,
			"nonce": req.Nonce,
		})
		if checkSign == "" || checkSign != req.Sign {
			ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "无效签名", "data": gin.H{}})
			ctx.Abort()
			return
		}

		// 创建 nonce
		_, err = ioc.SessionService().Create(context, req.AppID, req.Nonce, ctx.ClientIP())
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "内部错误", "data": gin.H{}})
			ctx.Abort()
			return
		}

		// 注入相关信息
		ctx.Set("product_id", productRow.ID)
		ctx.Next()
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/skyisboss/pay-system/internal/app/api/handler"
	"github.com/skyisboss/pay-system/internal/app/api/middleware"
	"github.com/skyisboss/pay-system/internal/ioc"
)

func NewApi(r *gin.Engine, ioc *ioc.Container) {
	handle := &handler.Handler{
		Ioc: ioc,
	}
	// v1 := r.Group("/api/v1/", middleware.ApiMiddle(ioc))
	v1 := r.Group("/api/v1/")
	{
		v1.POST("/withdraw", handle.ApplyWithdraw)
		v1.POST("/account", handle.CreateAccount)
	}
}
func NewAdmin(r *gin.Engine, ioc *ioc.Container) {
	handle := &handler.Handler{
		Ioc: ioc,
	}
	r.Use(middleware.CorsMiddleware())
	public := r.Group("/pk-admin")
	{
		public.GET("/login", handle.Login)
	}
	// v1 := r.Group("/admin/", middleware.ApiMiddle)
	// {
	// 	v1.POST("/address", handle.Address)
	// 	v1.POST("/withdraw", handle.Withdraw)
	// }
}

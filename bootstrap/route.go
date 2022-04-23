package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoute(ctx *gin.Engine) {
	//设置全局中间件
	registerGlobalMiddleWare(ctx)
	//注册api路由
	RegisterAPIRoutes(ctx)
	//配置404路由
}

//设置全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(gin.Logger(), gin.Recovery())
}

//注册网页路由
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"hellw": "word!",
			})
		})
	}
}

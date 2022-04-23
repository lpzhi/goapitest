package main

import (
	"fmt"
	"goapitest/bootstrap"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	bootstrap.SetupRoute(r)
	r.Use(gin.Logger(), gin.Recovery())

	r.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "页面 404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"errer_message": "路由未定义",
			})
		}
	})

	err := r.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

package main

import (
	"flag"
	"fmt"
	"goapitest/bootstrap"
	bstconfig "goapitest/config"
	"goapitest/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	bstconfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	bstconfig.Initialize()
	r := gin.New()
	bootstrap.SetupRoute(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Fatalln("test")
}

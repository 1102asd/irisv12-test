package main

import (
	"flag"
	"github.com/kataras/iris/v12"
	"project1/app/conf"
	"project1/app/route"
)

func main() {
	flag.Parse()
	app := iris.New()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+conf.MysqlConfig.ServerPort), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	app.AllowMethods(iris.MethodOptions)
	return app
}

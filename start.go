package main

import (
	"flag"
	"github.com/kataras/iris/v12"
	"irisv12-test/app/conf"
	"irisv12-test/app/route"
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

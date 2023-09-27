package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"irisv12-test/app/controllers"
	"irisv12-test/app/repo"
	"irisv12-test/app/service"
)

func InitRouter(app *iris.Application) {
	//app.Use(CrossAccess)
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/user")).Register(service.NewUserServices(repo.UserRepo)).Handle(new(controllers.UserController))
	//mvc.New(app.Party(bathUrl + "/book")).Handle(controllers.NewBookController())
	mvc.New(app.Party(bathUrl + "/book")).Register(service.NewBookServices(repo.BookRepo)).Handle(new(controllers.BookController))

}

package controllers

import (
	"github.com/kataras/iris/v12"
	"irisv12-test/app/models"
	"irisv12-test/app/service"
	"log"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func (g *UserController) PostLogin() models.Result {
	var m map[string]string
	log.Println("PostLogin Called")

	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
		return models.Result{Code: -1, Msg: "读取 JSON 错误"}
	}

	result := g.Service.Login(m)
	return result
}

func (g *UserController) PostSave() (result models.Result) {
	log.Println("PostSave Called")

	var user models.User
	if err := g.Ctx.ReadJSON(&user); err != nil {
		log.Println("ReadJSON Error:", err)
		result.Msg = "数据错误"
		return
	}

	return g.Service.Save(user)
}

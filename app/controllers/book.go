package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"irisv12-test/app/models"
	"irisv12-test/app/service"
)

type BookController struct {
	Ctx     iris.Context
	Service service.BookService
}

func (g *BookController) Get() (result models.Result) {
	m := ConvertStringMapToStringInterfaceMap(g.Ctx.URLParams())
	if m["page"] == "" || m["page"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 page"
		return
	}
	if cast.ToUint(m["page"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 page"
		return
	}
	if m["size"] == "" || m["size"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 size"
		return
	}
	if cast.ToUint(m["size"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 size"
		return
	}
	return g.Service.Get(m)
}

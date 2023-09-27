package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"irisv12-test/app/models"
	"irisv12-test/app/service"
	"log"
)

type BookController struct {
	Ctx     iris.Context
	Service service.BookService
}

func (g *BookController) Get() (result models.Result) {
	/**
	书籍查询接口 /api/book
	参数含有id时 查询详情
	否则为查询全部 根据page与size进行分页
	**/
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
func (g *BookController) GetBy(id uint) (result models.Result) {
	result, _ = g.Service.GetBookById(id)
	return
}
func (g *BookController) DeleteBy(id uint) (result models.Result) {
	g.Service.DelBookById(id)
	return
}
func (g *BookController) PutBy(id uint) (result models.Result) {
	_, book := g.Service.GetBookById(id)
	if book.ID == 0 {
		result.Msg = "未找到原始book数据"
		return
	}
	if err := g.Ctx.ReadJSON(&book); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}
	return g.Service.UpdateBookById(book)
}
func (g *BookController) Post() (result models.Result) {
	var book models.Book
	if err := g.Ctx.ReadJSON(&book); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}
	return g.Service.UpdateBookById(book)
}

package service

import (
	"irisv12-test/app/models"
	"irisv12-test/app/repo"
)

type BookService interface {
	Get(m map[string]interface{}) (result models.Result)
	GetBookById(id uint) (result models.Result, book models.Book)
	DelBookById(id uint) (result models.Result)
	UpdateBookById(book models.Book) (result models.Result)
}
type bookServices struct {
	bookRepo repo.BookRepository
}

func NewBookServices(repo repo.BookRepository) BookService {
	return &bookServices{bookRepo: repo}
}

func (b bookServices) Get(m map[string]interface{}) (result models.Result) {
	total, books := b.bookRepo.Get(m)
	maps := make(map[string]interface{}, 2)
	maps["total"] = total
	maps["list"] = books
	result.Data = maps
	result.Msg = "SUCCESS"
	result.Code = 0
	return
}
func (b bookServices) GetBookById(id uint) (result models.Result, book models.Book) {
	//TODO implement me
	book, err := b.bookRepo.GetBookById(id)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
	} else {
		result.Data = book
		result.Code = 0
		result.Msg = "SUCCESS"
	}
	return
}

func (b bookServices) DelBookById(id uint) (result models.Result) {
	err := b.bookRepo.DelBookById(id)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
	} else {
		result.Code = 0
		result.Msg = "SUCCESS"
	}
	return
}

func (b bookServices) UpdateBookById(book models.Book) (result models.Result) {
	err := b.bookRepo.UpdateBookById(book)
	if err != nil {
		result.Code = -1
		result.Msg = "保存失败"
	} else {
		result.Code = 1
		result.Msg = "保存成功"
	}
	return
}

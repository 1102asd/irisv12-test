package service

import (
	"irisv12-test/app/models"
	"irisv12-test/app/repo"
)

type BookService interface {
	Get(m map[string]interface{}) (result models.Result)
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

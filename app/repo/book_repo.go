package repo

import (
	"github.com/spf13/cast"
	"irisv12-test/app/datasource"
	"irisv12-test/app/models"
	"log"
)

type BookRepository interface {
	Get(m map[string]interface{}) (total int, books []models.Book)
}
type bookRepository struct {
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

var db = datasource.GetDB()
var BookRepo = NewBookRepository()

func (b bookRepository) Get(m map[string]interface{}) (total int, books []models.Book) {
	db.Table("book").Count(&total)
	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&books).Error
	if err != nil {
		log.Println("GetBookList Error:", err)
		panic("GetBookList Error")
	}
	return
}

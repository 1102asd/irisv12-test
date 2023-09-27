package repo

import (
	"github.com/spf13/cast"
	"irisv12-test/app/datasource"
	"irisv12-test/app/models"
	"log"
)

type BookRepository interface {
	Get(m map[string]interface{}) (total int, books []models.Book)
	GetBookById(id uint) (book models.Book, err error)
	DelBookById(id uint) (err error)
	UpdateBookById(book models.Book) (err error)
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
func (b bookRepository) GetBookById(id uint) (book models.Book, err error) {
	//TODO implement me
	err = db.First(&book, id).Error
	if err != nil {
		log.Println("GetBookById Error:", err)
	}
	return
}
func (b bookRepository) DelBookById(id uint) (err error) {
	var book models.Book
	book.ID = id
	err = db.Unscoped().Delete(&book).Error //如果直接Delete是软删除
	return
}
func (b bookRepository) UpdateBookById(book models.Book) (err error) {
	if book.ID != 0 {
		err := db.Save(&book).Error
		return err
	} else {
		err := db.Create(&book).Error
		return err
	}
}

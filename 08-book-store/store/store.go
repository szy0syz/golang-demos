package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	Id      string   `json:"id"`      // 图书ISBN
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   // 出版社
}

type Store interface {
	Create(*Book) error       //增
	Delete(string) error      //删
	Update(*Book) error       // 改
	Get(string) (Book, error) // 查
	GetAll() ([]Book, error)  // 查所有
}

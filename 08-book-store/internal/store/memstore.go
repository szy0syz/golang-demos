package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"fmt"
	"sync"
)

func init() {
	fmt.Println("注册mem")
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync  sync.RWMutex
	books map[string]*mystore.Book
}

func (m MemStore) Create(book *mystore.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Update(book *mystore.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Get(s string) (mystore.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) GetAll() ([]mystore.Book, error) {
	//TODO implement me
	panic("implement me")
}

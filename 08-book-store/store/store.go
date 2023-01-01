package store

type Book struct {
	Id      string   `json:"id"`   // 图书ISBN
	Name    string   `json:"name"` // 图书名称
	Authors []string `json:"authors"`
	Press   string
}

type Store interface {
	Create(*Book) error       //增
	Delete(string) error      //删
	Update(*Book) error       // 改
	Get(string) (Book, error) //查
	GetAll() ([]Book, error)  // 获取所有图书
}

package store

type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

type Store interface {
	Create(book *Book) error
	Update(book *Book) error
	Get(book string) (Book, error)
	GetAll() ([]Book, error)
	Delete(book string) error
}

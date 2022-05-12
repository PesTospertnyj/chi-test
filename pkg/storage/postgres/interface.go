package postgres

// Postgres ...
type Postgres interface {
	// Seed ...
	Seed() error
	// GetAllBooks ...
	GetAllBooks() ([]Book, error)
	// GetBooksByAuthor ...
	GetBooksByAuthor(string) ([]Book, error)
	// GetBooksByTitle ...
	GetBooksByTitle(string) ([]Book, error)
	// AddBook ...
	AddBook(*Book) error
	// UpdateBook ...
	UpdateBook(*Book) error
	// DeleteBook ...
	DeleteBook(int) error
}

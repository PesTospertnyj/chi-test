package postgres

import (
	"database/sql"
	"errors"
)

const (
	getBooksByAuthorQuery = `select id, title, author from books where author = $1;`
	getBooksByTitleQuery  = `select id, title, author from books where title = $1;`
	getAllBooksQuery      = `select id, title, author from books;`
	addBookQuery          = `insert into books (title, author) values ($1, $2);`
	updateBookQuery       = `update books set title = $1, author = $2 where id = $3;`
	deleteBookQuery       = `delete from books where id = $1;`
)

func (p *postgres) GetBooksByAuthor(author string) ([]Book, error) {
	rows, err := p.db.Query(getBooksByAuthorQuery, author)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return scanRows(rows)
}

func (p *postgres) GetBooksByTitle(title string) ([]Book, error) {
	rows, err := p.db.Query(getBooksByTitleQuery, title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return scanRows(rows)
}

func (p *postgres) GetAllBooks() ([]Book, error) {
	rows, err := p.db.Query(getAllBooksQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return scanRows(rows)
}

func (p *postgres) AddBook(book *Book) error {
	_, err := p.db.Exec(addBookQuery, book.Title, book.Author)

	return err
}

func (p *postgres) UpdateBook(book *Book) error {
	_, err := p.db.Exec(updateBookQuery, book.Title, book.Author, book.ID)

	return err
}

func (p *postgres) DeleteBook(id int) error {
	_, err := p.db.Exec(deleteBookQuery, id)

	return err
}

func scanRows(rows *sql.Rows) ([]Book, error) {
	books := make([]Book, 0)
	for rows.Next() {
		book := Book{}
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

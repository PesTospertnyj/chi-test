package util

import (
	"chi-test/pkg/dto"
	"chi-test/pkg/storage/postgres"
)

func BookDTOToBook(book *dto.Book) *postgres.Book {
	return &postgres.Book{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}
}

func BooksToBookDTOs(books []postgres.Book) []dto.Book {
	dtos := make([]dto.Book, 0)
	for _, book := range books {
		dtos = append(dtos, dto.Book{ID: book.ID, Title: book.Title, Author: book.Author})
	}

	return dtos
}

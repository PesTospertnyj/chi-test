package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"chi-test/pkg/dto"
	"chi-test/pkg/util"
)

func (a *api) GetBooks(c echo.Context) error {
	title := c.QueryParam("title")
	if title != "" {
		return a.getBooksByTitle(c, title)
	}

	author := c.QueryParam("author")
	if author != "" {
		return a.getBooksByAuthor(c, author)
	}

	return a.getAllBooks(c)
}

func (a *api) AddBook(c echo.Context) error {
	b := new(dto.Book)
	if err := c.Bind(b); err != nil {
		return err
	}

	book := util.BookDTOToBook(b)

	err := a.Storage.AddBook(book)
	if err != nil {
		a.Log.Errorf("error while adding new book: %v error: %e", book, err)

		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error while adding new book: %v error: %e", book, err))
	}

	return c.NoContent(http.StatusOK)
}

func (a *api) UpdateBook(c echo.Context) error {
	b := new(dto.Book)
	if err := c.Bind(b); err != nil {
		return err
	}

	book := util.BookDTOToBook(b)

	err := a.Storage.UpdateBook(book)
	if err != nil {
		a.Log.Errorf("error while updating book: %v error: %e", book, err)

		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error while updating  book: %v error: %e", book, err))
	}

	return c.NoContent(http.StatusOK)
}

func (a *api) DeleteBook(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		a.Log.Errorf("error while parsing param: %v error: %e", param, err)

		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error while parsing param: %v error: %e", param, err))
	}

	err = a.Storage.DeleteBook(id)
	if err != nil {
		a.Log.Errorf("error while deleting book with id: %d error: %e", id, err)

		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}

	return c.NoContent(http.StatusOK)
}

func (a *api) getBooksByTitle(c echo.Context, title string) error {
	books, err := a.Storage.GetBooksByTitle(title)
	if err != nil {
		a.Log.Errorf("error while getting books by title: %s; error: %e", title, err)

		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, util.BooksToBookDTOs(books))
}

func (a *api) getBooksByAuthor(c echo.Context, author string) error {
	books, err := a.Storage.GetBooksByAuthor(author)
	if err != nil {
		a.Log.Errorf("error while getting books by author: %s; error: %e", author, err)

		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, util.BooksToBookDTOs(books))
}

func (a *api) getAllBooks(c echo.Context) error {
	books, err := a.Storage.GetAllBooks()
	if err != nil {
		a.Log.Errorf("error while getting all books error: %e", err)

		return echo.NewHTTPError(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, util.BooksToBookDTOs(books))
}

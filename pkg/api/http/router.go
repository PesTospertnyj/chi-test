package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *api) NewRouter() *echo.Echo {
	e := echo.New()
	conf := middleware.DefaultLoggerConfig
	conf.Output = a.Log.Out
	e.Use(middleware.LoggerWithConfig(conf))
	e.Use(middleware.Recover())

	e.GET("/books", a.GetBooks)
	e.POST("/books", a.AddBook)
	e.PUT("/books", a.UpdateBook)
	e.DELETE("/books/:id", a.DeleteBook)

	return e
}

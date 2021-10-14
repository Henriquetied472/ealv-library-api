package server

import (
	"ealv-library-api/book"
	"ealv-library-api/db"
	"net/http"

	"github.com/henriquetied472/logplus"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echoServer *echo.Echo
}

func NewServer() *Server {
	return &Server{
		echoServer: echo.New(),
	}
}

func (s *Server) Start() {
	s.echoServer.GET("/books", func(c echo.Context) error {
		return c.String(http.StatusOK, db.ListBooksJSON())
	})
	s.echoServer.POST("/register", func(c echo.Context) error {
		form := make(map[string]string)
		form["title"] = c.Request().Form.Get("bTitle")
		form["autor"] = c.Request().Form.Get("bAutor")
		form["editor"] = c.Request().Form.Get("bEditor")
		form["pubYear"] = c.Request().Form.Get("bPubYear")
		form["dtLimit"] = c.Request().Form.Get("bDtLimit")
		form["sName"] = c.Request().Form.Get("bSName")
		form["sClass"] = c.Request().Form.Get("bSClass")
		form["raCod"] = c.Request().Form.Get("bSRaCod")
		form["raDig"] = c.Request().Form.Get("bSRaDig")
		form["raUf"] = c.Request().Form.Get("bSRaUf")

		b := book.BookFromForm(form)

		db.AddBook(b)

		return nil
	})
	s.echoServer.PUT("/edit", func(c echo.Context) error {
		aForm := make(map[string]string)
		bForm := make(map[string]string)

		aForm["title"] = c.Request().Form.Get("abTitle")
		aForm["autor"] = c.Request().Form.Get("abAutor")
		aForm["editor"] = c.Request().Form.Get("abEditor")
		aForm["pubYear"] = c.Request().Form.Get("abPubYear")
		aForm["dtLimit"] = c.Request().Form.Get("abDtLimit")
		aForm["sName"] = c.Request().Form.Get("abSName")
		aForm["sClass"] = c.Request().Form.Get("abSClass")
		aForm["raCod"] = c.Request().Form.Get("abSRaCod")
		aForm["raDig"] = c.Request().Form.Get("abSRaDig")
		aForm["raUf"] = c.Request().Form.Get("abSRaUf")

		bForm["title"] = c.Request().Form.Get("bbTitle")
		bForm["autor"] = c.Request().Form.Get("bbAutor")
		bForm["editor"] = c.Request().Form.Get("bbEditor")
		bForm["pubYear"] = c.Request().Form.Get("bbPubYear")
		bForm["dtLimit"] = c.Request().Form.Get("bbDtLimit")
		bForm["sName"] = c.Request().Form.Get("bbSName")
		bForm["sClass"] = c.Request().Form.Get("bbSClass")
		bForm["raCod"] = c.Request().Form.Get("bbSRaCod")
		bForm["raDig"] = c.Request().Form.Get("bbSRaDig")
		bForm["raUf"] = c.Request().Form.Get("bbSRaUf")

		ab := book.BookFromForm(aForm)
		bb := book.BookFromForm(bForm)

		db.EditBook(ab, bb)

		return nil
	})

	logplus.Fatal(s.echoServer.Start(":3030").Error())
}

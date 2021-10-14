package server

import (
	"ealv-library-api/book"
	"ealv-library-api/db"
	"net/http"

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
		return c.JSON(http.StatusOK, db.ListBooksJSON())
	})
	s.echoServer.POST("/register", func(c echo.Context) error {
		form := make(map[string]string)
		form["title"] = c.FormValue("bTitle")
		form["autor"] = c.FormValue("bAutor")
		form["editor"] = c.FormValue("bEditor")
		form["pubYear"] = c.FormValue("bPubYear")
		form["dtLimit"] = c.FormValue("bDtLimit")
		form["sName"] = c.FormValue("bSName")
		form["sClass"] = c.FormValue("bSClass")
		form["raCod"] = c.FormValue("bSRaCod")
		form["raDig"] = c.FormValue("bSRaDig")
		form["raUf"] = c.FormValue("bSRaUf")

		b := book.BookFromForm(form)

		db.AddBook(b)

		return nil
	})
	s.echoServer.PUT("/edit", func(c echo.Context) error {
		aForm := make(map[string]string)
		bForm := make(map[string]string)

		aForm["title"] = c.FormValue("abTitle")
		aForm["autor"] = c.FormValue("abAutor")
		aForm["editor"] = c.FormValue("abEditor")
		aForm["pubYear"] = c.FormValue("abPubYear")
		aForm["dtLimit"] = c.FormValue("abDtLimit")
		aForm["sName"] = c.FormValue("abSName")
		aForm["sClass"] = c.FormValue("abSClass")
		aForm["raCod"] = c.FormValue("abSRaCod")
		aForm["raDig"] = c.FormValue("abSRaDig")
		aForm["raUf"] = c.FormValue("abSRaUf")

		bForm["title"] = c.FormValue("bbTitle")
		bForm["autor"] = c.FormValue("bbAutor")
		bForm["editor"] = c.FormValue("bbEditor")
		bForm["pubYear"] = c.FormValue("bbPubYear")
		bForm["dtLimit"] = c.FormValue("bbDtLimit")
		bForm["sName"] = c.FormValue("bbSName")
		bForm["sClass"] = c.FormValue("bbSClass")
		bForm["raCod"] = c.FormValue("bbSRaCod")
		bForm["raDig"] = c.FormValue("bbSRaDig")
		bForm["raUf"] = c.FormValue("bbSRaUf")

		ab := book.BookFromForm(aForm)
		bb := book.BookFromForm(bForm)

		db.EditBook(ab, bb)

		return nil
	})
	s.echoServer.DELETE("/delete", func(c echo.Context) error {
		return nil
	})

	s.echoServer.Start(":3030")
}
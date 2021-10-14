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

		b := book.RegisterNewBook(form["title"], form["autor"], form["editor"], form["pubYear"])

		db.AddBook(*b)

		return nil
	})

	logplus.Fatal(s.echoServer.Start(":3030").Error())
}

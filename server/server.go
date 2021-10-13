package server

import (
	"net/http"
	"ealv-library-api/db"

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

	s.echoServer.Start(":3000")
}
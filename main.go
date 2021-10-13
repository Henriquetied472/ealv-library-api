package main

import (
	"ealv-library-api/server"
)

func main() {
	s := server.NewServer()
	
	s.Start()
}
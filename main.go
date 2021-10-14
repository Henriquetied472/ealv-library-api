package main

import (
	"ealv-library-api/dotenv-decoder"
	"ealv-library-api/server"
	"log"
)

func main() {
	err := dotenvdecoder.LoadENV()
	if err != nil {
		log.Fatalf("DEBUG: %v", err)
	}

	s := server.NewServer()
	
	s.Start()
}
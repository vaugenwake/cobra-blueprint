package api

import (
	"cobra-api/internal/server"
	"fmt"
)

func NewApi() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error: %v", err)
		panic("cannot start server")
	}
}

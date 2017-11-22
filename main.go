package main

import (
	"service"
)

func main() {
	s := server.NewServer()
	s.Run(":8080")
}

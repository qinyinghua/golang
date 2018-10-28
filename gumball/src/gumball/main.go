package main

import (
	"os"
)

func main() {

	port := os.Getenv("PORT")
	print(port)
	if len(port) < 4 {
		port = "3000"
	}

	server := NewServer()
	server.Run(":" + port)
}

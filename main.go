package main

import "os"

func main() {
	server := Server{}

	server.Initialize()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	server.Run(":" + port)
}

package main

import (
	"go-gameroom/driver/routing"
)

func main() {
	println("Server running...")
	// router := routing.NewHTTPRouter()
	router := routing.NewGinRouter()
	router.Run("localhost:8000")
}

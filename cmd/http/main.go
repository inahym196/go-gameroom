package main

import (
	"go-gameroom/controller"
)

func main() {
	mrouter := controller.GetMRouter()
	router := controller.GetRouter(mrouter)
	router.Run("localhost:3000")
}

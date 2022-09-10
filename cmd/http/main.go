package main

import (
	"fmt"
	"go-gameroom/controller"
)

func main() {
	fmt.Println("hoge")
	mrouter := controller.GetMRouter()
	router := controller.GetRouter(mrouter)
	router.Run("localhost:3000")
}

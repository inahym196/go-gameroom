package main

import (
	"go-gameroom/driver"
)

func main() {
	println("Server running...")
	driver.Serve("localhost:8000")
}

package main

import (
	"fmt"
	"go-starter/internal/database"
)

func main() {
	fmt.Println("Hello, World")

	database.InitDB()
}

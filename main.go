package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
		return
	}

	user := os.Getenv("USER")
	fmt.Println(user)
}

package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hi")
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintln("big problems!"))
	}

}
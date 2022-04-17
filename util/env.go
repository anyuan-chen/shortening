package util

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv (){
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintln(err))
	}
}
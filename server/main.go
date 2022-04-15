package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("hi")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

}
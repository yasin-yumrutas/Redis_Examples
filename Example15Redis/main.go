package main

import (
	"fmt"
	"log"
	"test/goP"
)

func main() {
	redisConfig := goP.RedisConfig{
		Adress:   "localhost",
		Password: "redispw",
		Port:     32768,
	}

	client, err := redisConfig.Connect()
	if err != nil {
		log.Fatalf("Error: redisConfig.Connect,  %v", err)
	}

	info, _ := client.Info()

	fmt.Println(info.Message)
}

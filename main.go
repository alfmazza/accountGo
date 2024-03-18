package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPosgresStore()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)

	server := NewAPIServer(":8083", store)
	server.Run()
}

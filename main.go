package main

import (
	"log"
)

func main() {
	store, err := NewPosgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8083", store)
	server.Run()
}

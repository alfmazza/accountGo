package main

func main() {
	server := NewAPIServer(":8083")
	server.Run()
}

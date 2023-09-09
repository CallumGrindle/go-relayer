package main

import (
	"api/internal/api"
	"api/internal/services"
	"fmt"
)

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)

	go services.Signer()

	api.StartServer(addr)
}

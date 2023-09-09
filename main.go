package main

import (
	"api/internal/api"
	"api/internal/config"
	"api/internal/services"
	"fmt"
)

func main() {
	config.InitConfig()

	addr := fmt.Sprintf(":%d", config.ApplicationConfig.Port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)

	go services.Signer()
	go services.Sender()
	go services.Listener()

	api.StartServer(addr)
}

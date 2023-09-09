package api

import (
	"net/http"
)

func StartServer(listenAddr string) {
	router := http.NewServeMux()
	initRoutes(router)

	err := http.ListenAndServe(listenAddr, router)

	if err != nil {
		panic(err)
	}
}

func initRoutes(router *http.ServeMux) {
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/execute", executeRequestHandler)
}

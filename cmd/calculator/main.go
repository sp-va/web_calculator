package main

import (
	"calculator/internal/middleware"
	"calculator/internal/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", middleware.ValidationMiddleware(routes.CalculateHandler))

	log.Println("запуск сервера на порту :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("ошибка: %s", err)
	}
}

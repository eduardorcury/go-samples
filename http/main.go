package http

import (
	"log"
	"net/http"
)

func main() {
	userRepository := NewUserRepository()
	service := NewUserService(userRepository)
	mux := http.NewServeMux()

	httpHandler := HttpHandler{
		Service: service,
	}

	mux.HandleFunc("POST /users", httpHandler.HandleNewUser)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}

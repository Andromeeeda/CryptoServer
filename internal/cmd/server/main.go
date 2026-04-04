package main

import (
	"cryptoserver/internal/repository"
	"cryptoserver/internal/service"
	httphandlers "cryptoserver/internal/transport/httpHandlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	userRep := repository.NewUserRep()
	authService := service.NewAuthService(userRep)
	authHandler := httphandlers.NewAuthHandler(authService)

	router.Path("/auth/register").Methods("POST").HandlerFunc(authHandler.Register)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("fail to listen server")
	}

}

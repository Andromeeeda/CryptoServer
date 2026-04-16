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

	cryptoRep := repository.NewCryptoRep()
	coinGeckoApiKey := "CG-Q3SdkQ3AZe2nPqELkZaivyBW"
	cryptoService := service.NewCryptoService(cryptoRep, coinGeckoApiKey)
	cryptoHandler := httphandlers.NewCryptoHandler(cryptoService)

	router.Path("/auth/register").Methods("POST").HandlerFunc(authHandler.Register)
	router.Path("/auth/login").Methods("POST").HandlerFunc(authHandler.Login)

	router.Path("/crypto").Methods("POST").HandlerFunc(cryptoHandler.CreateCrypto)
	router.Path("/crypto").Methods("GET").HandlerFunc(cryptoHandler.GetCrypto)
	router.Path("/crypto/{symbol}").Methods("GET").HandlerFunc(cryptoHandler.GetCryptoBySymbol)
	router.Path("/crypto/{symbol}/refresh").Methods("PUT").HandlerFunc(cryptoHandler.RefreshCryptoPrice)
	router.Path("/crypto/{symbol}/history").Methods("GET").HandlerFunc(cryptoHandler.GetCryptoHistory)
	router.Path("/crypto/{symbol}/stats").Methods("GET").HandlerFunc(cryptoHandler.GetCryptoStatistic)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("fail to listen server")
	}

}

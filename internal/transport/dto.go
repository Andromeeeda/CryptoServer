package transport

import (
	"cryptoserver/internal/repository"
	"encoding/json"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponce struct {
	JwtToken string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponce struct {
	JwtToken string `json:"token"`
}

type ErrorsDTO struct {
	Erorr string    `json:"error"`
	Time  time.Time `json:"time"`
}

func ErrorsDtoToString(errDTO *ErrorsDTO) string {
	errDTOJson, _ := json.Marshal(errDTO)

	return string(errDTOJson)
}

type CryptoSymbolReguest struct {
	Symbol string `json:"symbol"`
}

type CryptoResponce struct {
	Symbol        string    `json:"symbol"`
	Name          string    `json:"name"`
	Current_price float64   `json:"current_price"`
	Last_updated  time.Time `json:"last_updated"`
}

type GetCryptosResponce struct {
	Coins []*repository.Crypto `json:"cryptos"`
}

type CryptoHistoryResponce struct {
	Symbol        string    `json:"symbol"`
	History 	[]repository.PriceEntry `json:"history"`
}

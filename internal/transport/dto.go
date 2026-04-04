package transport

import (
	"encoding/json"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ErrorsDTO struct {
	Erorr string    `json:"error"`
	Time  time.Time `json:"time"`
}

func ErrorsDtoToString(errDTO *ErrorsDTO) string {
	errDTOJson, _ := json.Marshal(errDTO)

	return string(errDTOJson)
}

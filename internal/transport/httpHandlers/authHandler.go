package httphandlers

import (
	"cryptoserver/internal/service"
	"cryptoserver/internal/transport"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest transport.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		errDto := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)

		http.Error(w, transport.ErrorsDtoToString(&errDto), http.StatusBadRequest)

		return
	}

	//далее передать данные которые преобразовали из JSon в service для создания пользователя
	//Создать пользователя и обработать ошибки
	//Вернуть тело ответа и статус код
	if err := h.authService.RegisterUser(registerRequest.Username, registerRequest.Password); err != nil {
		errDto := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}
		fmt.Println("Error!", errDto)

		http.Error(w, transport.ErrorsDtoToString(&errDto), http.StatusConflict)

	}

	send := "Пользователь зарегестрирован"
	fmt.Println(send)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(send))

}

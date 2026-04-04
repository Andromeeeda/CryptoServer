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
	jwtToken, err := h.authService.RegisterUser(registerRequest.Username, registerRequest.Password)
	if err != nil {
		errDto := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}
		fmt.Println("Error!", errDto)

		http.Error(w, transport.ErrorsDtoToString(&errDto), http.StatusConflict)

		return
	}

	responce := transport.RegisterResponce{
		JwtToken: jwtToken,
	}

	msg := "The user is registered"
	fmt.Println(msg)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responce)

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	//1.Считать данные пользователя
	//2. Выполнить бизнес-логику service
	//3. Вернуть ответ

	var loginRequest transport.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusBadRequest)

		return
	}

	jwtToken, err := h.authService.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		errDto := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}
		fmt.Println("Error!", errDto)

		http.Error(w, transport.ErrorsDtoToString(&errDto), http.StatusUnauthorized)

		return
	}

	responce := transport.LoginResponce{
		JwtToken: jwtToken,
	}

	fmt.Println("You have successfully logged into your account!")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responce)

}

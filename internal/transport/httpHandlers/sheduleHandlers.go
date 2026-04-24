package httphandlers

import (
	"cryptoserver/internal/middleware"
	"cryptoserver/internal/service"
	"cryptoserver/internal/transport"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type SheduleHandlers struct {
	sheduleService *service.SheduleService
}

func NewSheduleHandlers(sheduleService *service.SheduleService) *SheduleHandlers {
	return &SheduleHandlers{
		sheduleService: sheduleService,
	}
}

func (h *SheduleHandlers) GetShedule(w http.ResponseWriter, r *http.Request) {

	//Проверка аутентификации
	username, err := middleware.VerifyAuth(r)
	if err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("authentification error",errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusUnauthorized)

		return
	}

	log.Printf("User %s requested crypto list", username)


	sheduleInfo := h.sheduleService.GetSheduleInfo()

	var sheduleReponse = transport.SheduleResponce{
		Enabled:          sheduleInfo.Enabled,
		Interval_seconds: sheduleInfo.Interval_seconds,
		Last_update:      sheduleInfo.Last_update,
		Next_update:      sheduleInfo.Next_update,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sheduleReponse)

}

func (h *SheduleHandlers) PutShedule(w http.ResponseWriter, r *http.Request) {

	//Проверка аутентификации
	username, err := middleware.VerifyAuth(r)
	if err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("authentification error",errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusUnauthorized)

		return
	}

	log.Printf("User %s requested crypto list", username)


	var sheduleInfo transport.ChangeSheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&sheduleInfo); err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusBadRequest)

		return
	}

	shedule, err := h.sheduleService.ChangeShedule(&sheduleInfo)
	if err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusInternalServerError)

		return
	}

	sheduleResponce := transport.ChangeSheduleResponce{
		Enabled:          shedule.Enabled,
		Interval_seconds: shedule.Interval_seconds,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sheduleResponce)
}

func (h *SheduleHandlers) SheduleTrigger(w http.ResponseWriter, r *http.Request) {

	//Проверка аутентификации
	username, err := middleware.VerifyAuth(r)
	if err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("authentification error",errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusUnauthorized)

		return
	}

	log.Printf("User %s requested crypto list", username)

	
	update_count, timestamp, err := h.sheduleService.TriggerUpdate()
	if err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusInternalServerError)

		return
	}

	sheduleTriggerResponce := transport.TriggerSheduleResponce{
		Update_count: update_count,
		Timestamp:    timestamp,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sheduleTriggerResponce)
}

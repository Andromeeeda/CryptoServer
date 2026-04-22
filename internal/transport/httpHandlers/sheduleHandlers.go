package httphandlers

import (
	"cryptoserver/internal/service"
	"cryptoserver/internal/transport"
	"encoding/json"
	"fmt"
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

	var sheduleInfo transport.ChangeSheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&sheduleInfo); err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDTO)

		http.Error(w,transport.ErrorsDtoToString(&errDTO),http.StatusBadRequest)

		return
	}

	shedule,err :=  h.sheduleService.ChangeShedule(&sheduleInfo)
	if err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDTO)

		http.Error(w,transport.ErrorsDtoToString(&errDTO),http.StatusInternalServerError)

		return 
	}

	sheduleResponce := transport.ChangeSheduleResponce {
		Enabled: shedule.Enabled,
		Interval_seconds: shedule.Interval_seconds,
	}

	w.WriteHeader(http.StatusOK) 
	json.NewEncoder(w).Encode(sheduleResponce)
}


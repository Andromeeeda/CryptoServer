package httphandlers

import (
	"cryptoserver/internal/core"
	"cryptoserver/internal/service"
	"cryptoserver/internal/transport"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type CryptoHandler struct {
	cryptoService *service.CryptoService
}

func NewCryptoHandler(cryptoService *service.CryptoService) *CryptoHandler {
	return &CryptoHandler{
		cryptoService: cryptoService,
	}
}

func (c *CryptoHandler) CreateCrypto(w http.ResponseWriter, r *http.Request) {
	//В хенделере я буду получать в теле запроса JSon и его переформатировать
	// потом отдавать этот symbol на cryptoservice в котором будет логика запроса к
	// апишке GoinGecko и получение конкртеной криптовалюты тоесть м
	// не нужно будет внутри сервиса сделать http клиент и он у меня должен получить
	// все информацию о конкретной криптовалюте далее я эти данные криптовалюты отправляю
	// на хранение в repository. После чего отдаю ответ клиенту в виде JSon с криптовалютой

	var cryptosymbolRequest transport.CryptoSymbolReguest
	if err := json.NewDecoder(r.Body).Decode(&cryptosymbolRequest); err != nil {
		errDTO := transport.ErrorsDTO{
			Erorr: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDTO)

		http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusBadRequest)

		return
	}

	CryptoInfo, err := c.cryptoService.AddCrypto(cryptosymbolRequest.Symbol)
	if err != nil {

		if errors.Is(err, core.ErrBadRequest) {
			errDTO := transport.ErrorsDTO{
				Erorr: err.Error(),
				Time:  time.Now(),
			}

			fmt.Println("Error", errDTO)

			http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusBadRequest)
			return
		}

		if errors.Is(err, core.ErrConflict) {
			errDTO := transport.ErrorsDTO{
				Erorr: err.Error(),
				Time:  time.Now(),
			}

			fmt.Println("Error", errDTO)

			http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusConflict)
			return
		}

		if errors.Is(err, core.ErrInternalServerError) {
			errDTO := transport.ErrorsDTO{
				Erorr: err.Error(),
				Time:  time.Now(),
			}

			fmt.Println("Error", errDTO)

			http.Error(w, transport.ErrorsDtoToString(&errDTO), http.StatusInternalServerError)
			return
		}

	}

	cryptoResponce := transport.CryptoResponce{
		Symbol:        CryptoInfo.Symbol,
		Name:          CryptoInfo.Name,
		Current_price: CryptoInfo.Current_price,
		Last_updated:  CryptoInfo.Last_updated,
	}

	fmt.Println("Cryptocurrency successfully added")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cryptoResponce)

}

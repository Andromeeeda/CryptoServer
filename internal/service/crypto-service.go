package service

import (
	"cryptoserver/internal/core"
	"cryptoserver/internal/repository"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type CryptoService struct {
	cryptoRepository *repository.CryptoRep
	coinGeckoClient  *http.Client
	coinGeckoApiKey  string
}

func NewCryptoService(cryptoRepository *repository.CryptoRep, coinGeckoApiKey string) *CryptoService {
	return &CryptoService{
		cryptoRepository: cryptoRepository,
		coinGeckoClient:  &http.Client{Timeout: 10 * time.Second},
		coinGeckoApiKey: coinGeckoApiKey,
	}
}

func (s *CryptoService) AddCrypto(symbol string) (*repository.Crypto, error) {

	if symbol == "" {
		return nil, core.ErrBadRequest
	}

	coinId, err := s.getCoinId(symbol)
	if err != nil {
		return nil, err
	}

	coinInfo, err := s.getCoinInfo(coinId)
	if err != nil {
		return nil, err
	}

	coinPrice, err := s.getCurrentPrice(coinId)
	if err != nil {
		return nil, err
	}

	Crypto := repository.Crypto{
		Symbol:        coinInfo.Symbol,
		Name:          coinInfo.Name,
		Current_price: coinPrice,
		Last_updated:  coinInfo.Last_updated,
	}

	if err := s.cryptoRepository.AddCryptoRep(&Crypto); err != nil {
		return nil, err
	}

	return &Crypto, nil

}

// Получение CoinGeckoId по symbol
func (s *CryptoService) getCoinId(symbol string) (string, error) {

	url := "https://api.coingecko.com/api/v3/coins/list"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("x-cg-demo-api-key", s.coinGeckoApiKey)

	resp, err := s.coinGeckoClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	//Читаем и парсим JSON ответ
	var coins []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	}

	json.NewDecoder(resp.Body).Decode(&coins)

	for _, coin := range coins {
		if strings.ToUpper(coin.Symbol) == symbol {
			return coin.ID, nil
		}
	}

	return "", errors.New("Search error")

}

//Получение информации о криптовалюте 
func (s *CryptoService) getCoinInfo(coinID string) (*CoinInfo, error) {

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", coinID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-cg-demo-api-key", s.coinGeckoApiKey)

	resp, err := s.coinGeckoClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var CryproStruct CoinInfo

	if err := json.NewDecoder(resp.Body).Decode(&CryproStruct); err != nil {
		return nil, err
	}

	return &CryproStruct, nil

}

//Получение цены криптовалюты 
func (s *CryptoService) getCurrentPrice(coinID string) (float64, error) {

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", coinID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("x-cg-demo-api-key", s.coinGeckoApiKey)

	resp, err := s.coinGeckoClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	price, ok := result[coinID]["usd"]
	if !ok {
		return 0, errors.New(coinID + " Price not found")
	}

	return price, nil

}

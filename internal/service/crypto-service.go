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
		coinGeckoApiKey:  coinGeckoApiKey,
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
		Symbol:        strings.ToUpper(coinInfo.Symbol),
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

// Получение информации о криптовалюте
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

// Получение цены криптовалюты
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

func (s *CryptoService) ListCrypto() ([]*repository.Crypto, error) {

	return s.cryptoRepository.GetAllCryptos()
}

func (s *CryptoService) ListCryptoBySymbol(symbol string) (*repository.Crypto, error) {

	return s.cryptoRepository.GetCrypto(symbol)
}

func (s *CryptoService) RefreshPrice(symbol string) (*repository.Crypto, error) {

	_, err := s.cryptoRepository.GetCrypto(symbol)
	if err != nil {
		return nil, core.ErrNotFound
	}

	coinId, err := s.getCoinId(symbol)
	if err != nil {
		return nil, core.ErrInternalServerError
	}

	NewCurrentPrice, err := s.getCurrentPrice(coinId)
	if err != nil {
		return nil, core.ErrInternalServerError
	}

	Crypto, err := s.cryptoRepository.UpdatePrice(symbol, NewCurrentPrice)
	if err != nil {
		return nil, core.ErrInternalServerError
	}

	return Crypto, nil

}

func (s *CryptoService) RefreshAllPrices() (int,error) {

	cryptos,err := s.cryptoRepository.GetAllCryptos()
	if err != nil {
		return 0,err
	}

	if len(cryptos) == 0 {
		return 0,errors.New("No cryptocurrency added")
	}

	updatedcount := 0
	for _,crypto := range cryptos {
		if _,err := s.RefreshPrice(crypto.Symbol); err != nil {
			continue
		}
		updatedcount++
	}

	return updatedcount,nil 

}

func (s *CryptoService) CryptoHistory(symbol string) ([]repository.PriceEntry, error) {

	coinId, err := s.getCoinId(symbol)
	if err != nil {
		return nil, core.ErrInternalServerError
	}

	historyPrice, err := s.getCryptoPriceHistory(coinId)
	if err != nil {
		return nil, err
	}

	history := make([]repository.PriceEntry, 0, len(historyPrice))

	for _, point := range historyPrice {

		entry := repository.PriceEntry{
			Price: point[1],
			Time:  time.UnixMilli(int64(point[0])),
		}

		history = append(history, entry)
	}

	s.cryptoRepository.AddCryptoHistoryPrice(symbol, history)

	return history, nil
}

func (s *CryptoService) getCryptoPriceHistory(coinID string) ([][]float64, error) {

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=usd&days=30", coinID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, core.ErrInternalServerError
	}

	req.Header.Set("x-cg-demo-api-key", s.coinGeckoApiKey)

	resp, err := s.coinGeckoClient.Do(req)
	if err != nil {
		return nil, core.ErrInternalServerError
	}

	defer resp.Body.Close()

	var resultPrice struct {
		Prices [][]float64 `json:"prices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&resultPrice); err != nil {
		return nil, core.ErrInternalServerError
	}

	return resultPrice.Prices, nil

}

func (s *CryptoService) CryptoStatistic(symbol string) (*repository.Statistic, float64, error) {

	coinId, err := s.getCoinId(symbol)
	if err != nil {
		return nil, 0, core.ErrInternalServerError
	}

	prices, err := s.getCryptoPriceHistory(coinId)
	if err != nil {
		return nil, 0, err
	}

	if len(prices) == 0 {
		return nil, 0, errors.New("no price data")
	}

	// Извлекаем только цены из массива [timestamp, price]
	priceValues := make([]float64, len(prices))
	for i, point := range prices {
		priceValues[i] = point[1]
	}

	// Рассчет статистики
	minPrice := priceValues[0]
	maxPrice := priceValues[0]
	sum := 0.0

	for _, price := range priceValues {
		if price < minPrice {
			minPrice = price
		}
		if price > maxPrice {
			maxPrice = price
		}
		sum += price
	}

	avgPrice := sum / float64(len(priceValues))
	firstPrice := priceValues[0]
	lastPrice := priceValues[len(priceValues)-1]
	priceChange := lastPrice - firstPrice
	priceChangePercent := (priceChange / firstPrice) * 100

	currentPrice := lastPrice

	stats := &repository.Statistic{
		Min_price:            minPrice,
		Max_price:            maxPrice,
		Avg_price:            avgPrice,
		Price_change:         priceChange,
		Price_change_percent: priceChangePercent,
		Records_count:        len(priceValues),
	}

	//добавление на хранение
	s.cryptoRepository.AddCryptoStatistic(symbol, stats)

	return stats, currentPrice, nil

}

func (s *CryptoService) DeleteCryptoInfo(symbol string) error {

	_, err := s.cryptoRepository.GetCrypto(symbol)
	if err != nil {
		return err
	}

	s.cryptoRepository.DeleteCrypto(symbol)

	s.cryptoRepository.DeleteHistory(symbol)

	return nil
}

package repository

import (
	"errors"
)

type CryptoRep struct {
	coins   map[string]*Crypto
	history *CryptoHistoryPrice
	stats map[string]*Statistic
}

func NewCryptoRep() *CryptoRep {
	return &CryptoRep{
		coins: make(map[string]*Crypto),
		history: &CryptoHistoryPrice{
			History: make(map[string][]PriceEntry),
		},
		stats: make(map[string]*Statistic),
	}
}

func (r *CryptoRep) AddCryptoRep(crypto *Crypto) error {

	if _, exist := r.coins[crypto.Symbol]; exist {
		return errors.New("crypto already exists")
	}

	r.coins[crypto.Symbol] = crypto

	return nil

}

func (r *CryptoRep) GetAllCryptos() ([]*Crypto, error) {

	cryptos := make([]*Crypto, 0, len(r.coins))

	for _, val := range r.coins {
		cryptos = append(cryptos, val)
	}

	return cryptos, nil
}

func (r *CryptoRep) GetCrypto(symbol string) (*Crypto, error) {

	crypto, ok := r.coins[symbol]
	if !ok {
		return nil, errors.New("Crypto not found")
	}

	return crypto, nil
}

func (r *CryptoRep) UpdatePrice(symbol string, NewCurrentPrice float64) (*Crypto, error) {

	crypto, ok := r.coins[symbol]
	if !ok {
		return nil, errors.New("Crypto not found")
	}

	crypto.Current_price = NewCurrentPrice

	return crypto, nil

}

func (r *CryptoRep) AddCryptoHistoryPrice(symbol string, historyPrice []PriceEntry){

	r.history.History[symbol] = historyPrice

}

func (r *CryptoRep) AddCryptoStatistic(symbol string,stats *Statistic) {
	r.stats[symbol] = stats
}

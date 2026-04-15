package repository

import "errors"

type CryptoRep struct {
	coins map[string]*Crypto
}

func NewCryptoRep() *CryptoRep {
	return &CryptoRep{
		coins: make(map[string]*Crypto),
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

	return crypto,nil
}

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

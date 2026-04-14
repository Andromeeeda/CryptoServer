package service

import "time"

type CoinInfo struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Last_updated time.Time `json:"last_updated"`
}

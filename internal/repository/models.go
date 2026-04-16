package repository

import "time"

type User struct {
	Username string
	Password string
}

type Crypto struct {
	Symbol        string
	Name          string
	Current_price float64
	Last_updated  time.Time
}

type PriceEntry struct {
	Price float64
	Time  time.Time
}

type CryptoHistoryPrice struct {
	History map[string][]PriceEntry
}

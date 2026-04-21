package repository

import "time"

type User struct {
	Username string
	HashPassword string
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

type Statistic struct {
	Min_price            float64
	Max_price            float64
	Avg_price            float64
	Price_change         float64
	Price_change_percent float64
	Records_count        int
}

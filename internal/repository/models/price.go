package models

type Price struct {
	ID       int64   `json:"id"`
	Price    float64 `json:"price"`
	Wallet   string  `json:"wallet"`
	Exchange string  `json:"exchange"`
}

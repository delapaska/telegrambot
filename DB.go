package main

import "fmt"

type bnResp struct {
	Price float64 `json:"price,string"`
	Code  int64   `json:"code"`
}

type wallet map[string]float64

var db = map[int64]wallet{}

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "S8859306s"
	dbname   = "DB"
)

var dbInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type bnResp struct {
	Price float64 `json:"price,string"`
	Code  int64   `json:"code"`
}

func getPrice(symbol string) (price float64, err error) {
	resp, err := http.Get(fmt.Sprintf(os.Getenv("BINANCE_URL"), symbol))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var jsonResp bnResp

	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return
	}

	if jsonResp.Code != 0 {
		err = errors.New("wrong symbol")
	}
	price = jsonResp.Price

	return
}

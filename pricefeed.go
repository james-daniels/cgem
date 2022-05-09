package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	PRICEFEEDENDPOINT = "/v1/pricefeed"
)

type NewPrice struct {
	Pair                string `json:"pair"`
	Price               string `json:"price"`
	PercentageChange24h string `json:"percentChange24h"`
}

func GetPrice(symbol, baseurl string) *NewPrice {

	url := baseurl + PRICEFEEDENDPOINT

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get ecountered an error: %v", err)
	}
	defer resp.Body.Close()

	var np []NewPrice
	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&np)
		if err != nil {
			fmt.Printf("json new decoder ecountered an error: %v", err)
		}
	} else {
		resp.Body.Close()
		fmt.Printf("%v: ecountered an error: %v", resp.StatusCode, err)
	}

	for _, v := range np {
		if v.Pair == strings.ToUpper(symbol) {
			return &NewPrice{
				Pair:                v.Pair,
				Price:               v.Price,
				PercentageChange24h: v.PercentageChange24h,
			}
		}
	}
	return nil
}

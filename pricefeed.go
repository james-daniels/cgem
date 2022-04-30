package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type NewPrice struct {
	Pair string `json:"pair"`
	Price string `json:"price"`
	PercentageChange24 string `json:"percentChange24h"`
}

func PriceFeed(symbol, baseurl string, offset float64) string {
	endpoint := "/v1/pricefeed"
	url := baseurl + endpoint

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(fmt.Errorf("http.get ecountered an error: %v", err))
		return ""
	}
	defer resp.Body.Close()

	var np []NewPrice
	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&np)
		if err != nil {
		fmt.Println(fmt.Errorf("json decoder ecountered an error: %v", err))
		return ""
		}
	}
	var p float64
	for _, v := range np {
		if v.Pair == strings.ToUpper(symbol) {
			p, _ = strconv.ParseFloat(v.Price, 64)
		}
	}

	price := fmt.Sprint(p + offset)
	return price
}
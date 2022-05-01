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
	PercentageChange24h string `json:"percentChange24h"`
}

func priceFeed(symbol, baseurl string, offset float64) (string, error) {
	endpoint := "/v1/pricefeed"
	url := baseurl + endpoint

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("priceFeed: http.get ecountered an error: %v", err)
	}
	defer resp.Body.Close()

	var np []NewPrice
	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&np)
		if err != nil {
			return "", fmt.Errorf("priceFeed: json decoder ecountered an error: %v", err)
		}
	}
	var price float64
	for _, v := range np {
		if v.Pair == strings.ToUpper(symbol) {
			price, _ = strconv.ParseFloat(v.Price, 64)
		}
	}

	return fmt.Sprint(price + offset), nil
}
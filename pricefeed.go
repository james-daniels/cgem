package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type NewPrice struct {
	Pair                string `json:"pair"`
	Price               string `json:"price"`
	PercentageChange24h string `json:"percentChange24h"`
}

func PriceFeed(symbol, baseurl string, offset int) (string, error) {
	endpoint := "/v1/pricefeed"
	url := baseurl + endpoint

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http get ecountered an error: %v", err)
	}
	defer resp.Body.Close()

	var np []NewPrice
	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&np)
		if err != nil {
			return "", fmt.Errorf("json new decoder ecountered an error: %v", err)
		}
	} else {
		resp.Body.Close()
		return "", fmt.Errorf("%v: ecountered an error: %v", resp.StatusCode, err)
	}

	var price float64
	for _, v := range np {
		if v.Pair == strings.ToUpper(symbol) {
			price, err = strconv.ParseFloat(v.Price, 64)
			if err != nil {
				return "", fmt.Errorf("string convert parse float ecountered an error: %v", err)
			}
		}
	}

	return fmt.Sprint(price + float64(offset)), nil
}

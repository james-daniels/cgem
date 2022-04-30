package main

import (
	"flag"
	"fmt"
	"log"
)

var symbol string
var amount string
var price string
var side string
var env string

func init() {
	flag.StringVar(&symbol, "s", "", "SYMBOL: The symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: Quoted decimal amount to purchase")
	flag.StringVar(&price, "p", "", "PRICE: Quoted decimal amount to spend per unit")
	flag.StringVar(&side, "t", "buy", "TYPE: buy or sell")
	flag.StringVar(&env, "e", "sand", "ENVIRONMENT: prod or sand")
}

func main() {
	flag.Parse()

	var baseurl string

	switch env {
	case "prod":
		baseurl = "https://api.gemini.com"
	case "sand":
		baseurl = "https://api.sandbox.gemini.com"
	default:
		fmt.Println(`enter a value of either "prod" or "sand".`)
	}

	payload, err := PayloadBuilder(symbol, amount, price, side)
	if err != nil {
		fmt.Print(fmt.Errorf("ecountered an error: %v", err))
		return
	}

	signature := SigBuilder(payload)

	result, err := PostOrder(baseurl, payload, signature)
	if err != nil {
		fmt.Print(fmt.Errorf("ecountered an error: %v", err))
		return
	}

	log.Printf("%+v\n", result)
}

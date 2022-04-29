package main

import (
	"flag"
	"fmt"
)

var symbol string
var amount string
var price string
var side string
var env string

func init(){
	flag.StringVar(&symbol, "s", "", "SYMBOL: The symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: Quoted decimal amount to purchase")
	flag.StringVar(&price, "p", "", "PRICE: Quoted decimal amount to spend per unit")
	flag.StringVar(&side, "t", "buy", "TYPE: buy or sell")
	flag.StringVar(&env, "e", "sand", "ENVIRONMENT: prod or sand")
}


func main() {
	flag.Parse()

	var url string

	switch env {
	case "prod":
		url = "https://api.gemini.com"
	case "sand":
		url = "https://api.sandbox.gemini.com"
	}

	//payload, _ := PayloadBuilder("ltcusd", "1", "10.00", "buy")
	payload, _ := PayloadBuilder(symbol, amount, price, side)

	signature := SigBuilder(payload)

	result, _ := PostOrder(url, payload, signature)

	fmt.Printf("%+v\n", result)
}

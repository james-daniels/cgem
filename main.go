package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var symbol string
var amount string
var offset float64
var side string
var env string
var reoccur int

func init() {
	flag.StringVar(&symbol, "s", "", "SYMBOL: The symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: Quoted decimal amount to purchase")
	flag.Float64Var(&offset, "o", 0, "PRICE OFFSET: Quoted decimal amount to ADD TO PRICE")
	flag.StringVar(&side, "t", "buy", "TYPE: buy or sell")
	flag.StringVar(&env, "e", "sand", "ENVIRONMENT: prod or sand")
	flag.IntVar(&reoccur, "r", 0, `REOCCUR: frequency in hours of reocurrence (default "0")`)
}

func main() {
	flag.Parse()

	var baseurl string

	switch env {
	case "prod":
		baseurl = "https://api.gemini.com"
	case "sand":
		baseurl = "https://api.sandbox.gemini.co"
	default:
		fmt.Println(`enter a value of either "prod" or "sand".`)
	}

	for {
		price, err := priceFeed(symbol, baseurl, offset)
		errorHandler(err)

		payload, err := payloadBuilder(symbol, amount, price, side)
		errorHandler(err)

		signature := sigBuilder(payload)

		response, err := newOrder(baseurl, payload, signature)
		errorHandler(err)

		log.Printf("%+v\n", response)

		if reoccur <= 0 {
			return
		} else {
			time.Sleep(time.Hour * time.Duration(reoccur))
		}
	}
}


func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

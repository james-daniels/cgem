package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var symbol string
var amount string
var offset int
var side string
var env string
var reoccur int

func init() {
	flag.StringVar(&symbol, "s", "", "SYMBOL: symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: amount to purchase")
	flag.IntVar(&offset, "o", 0, `PRICE OFFSET: amount to ADD TO PRICE (default "0")`)
	flag.StringVar(&side, "t", "buy", "SIDE TYPE: buy or sell")
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
		baseurl = "https://api.sandbox.gemini.com"
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
		log.Fatalln(err)
	}
}

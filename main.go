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
var repeat int

func init() {
	flag.StringVar(&symbol, "s", "", "SYMBOL: symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: amount to purchase")
	flag.IntVar(&offset, "o", 0, `OFFSET: amount to ADD TO PRICE (default "0")`)
	flag.StringVar(&side, "t", "buy", "SIDE TYPE: buy or sell")
	flag.StringVar(&env, "e", "sand", "ENVIRONMENT: prod or sand")
	flag.IntVar(&repeat, "r", 0, `REPEAT: frequency in hours to repeat (default "0")`)
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
		errHandler(err)

		payload, err := payloadBuilder(symbol, amount, price, side)
		errHandler(err)

		signature := sigBuilder(payload)

		response, err := newOrder(baseurl, payload, signature)
		errHandler(err)

		log.Printf("%+v\n", response)

		if repeat <= 0 {
			return
		} else {
			fmt.Println()
			time.Sleep(time.Hour * time.Duration(repeat))
		}
	}
}

func errHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

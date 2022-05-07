package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	symbol string
	amount string
	offset int
	side   string
	env    string
	repeat int
)

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	env = cfg.Section("").Key("environment").String()
}

func init() {
	flag.StringVar(&symbol, "s", "", "SYMBOL: symbol for the new order")
	flag.StringVar(&symbol, "symbol", "", "SYMBOL: symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: amount to purchase")
	flag.StringVar(&amount, "amount", "", "AMOUNT: amount to purchase")
	flag.IntVar(&offset, "o", 0, `OFFSET: amount to ADD TO PRICE (default "0")`)
	flag.IntVar(&offset, "offset", 0, `OFFSET: amount to ADD TO PRICE (default "0")`)
	flag.StringVar(&side, "S", "buy", "SIDE TYPE: buy or sell")
	flag.StringVar(&side, "side", "buy", "SIDE TYPE: buy or sell")
	flag.IntVar(&repeat, "r", 0, `REPEAT: frequency in hours to repeat (default "0")`)
	flag.IntVar(&repeat, "repeat", 0, `REPEAT: frequency in hours to repeat (default "0")`)
}

func main() {
	flag.Parse()

	var baseurl string

	switch env {
	case "production":
		baseurl = "https://api.gemini.com"
	case "sandbox":
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

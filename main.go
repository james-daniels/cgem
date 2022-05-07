package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

var (
	symbol string
	amount string
	offset int
	side   string
	env    string
	repeat bool
)

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	env = cfg.Section("").Key("environment").String()
	repeat, err = cfg.Section("recurrence").Key("repeat").Bool()
	if err != nil {
		log.Fatalf("Failed to read parameter: %v", err)
	}
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
		fmt.Println(`enter a value of either "production" or "sandbox".`)
	}

	switch repeat {
	case true:
		MultiInst(baseurl)
	default:
		OneInst(baseurl)
	}
}

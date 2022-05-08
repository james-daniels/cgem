package main

import (
	"log"

	"time"

	"gopkg.in/ini.v1"
)

var (
	pretty bool
)

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	pretty, err = cfg.Section("").Key("pretty").Bool()
	if err != nil {
		log.Fatalf("Failed to read parameter: %v", err)
	}
}

func OneInst(baseurl string) {

	price, err := PriceFeed(symbol, baseurl, offset)
	errHandler(err)

	payload, err := PayloadBuilder(symbol, amount, price, side)
	errHandler(err)

	signature := SigBuilder(payload)

	response, err := NewOrder(baseurl, payload, signature)
	errHandler(err)

	if pretty {
		MakePretty(response)
	} else {
		log.Printf("%+v\n", response)
	}
}

func MultiInst(baseurl string, freq int) {

	if freq <= 0 {
		log.Fatalln("enter frequency value greater than 0")

	} else {

		for {
			price, err := PriceFeed(symbol, baseurl, offset)
			errHandler(err)

			payload, err := PayloadBuilder(symbol, amount, price, side)
			errHandler(err)

			signature := SigBuilder(payload)

			response, err := NewOrder(baseurl, payload, signature)
			errHandler(err)

			log.Printf("%+v\n\n", response)

			time.Sleep(time.Hour * time.Duration(freq))
		}
	}
}

func errHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

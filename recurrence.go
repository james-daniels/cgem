package main

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	freq int
)

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	freq, err = cfg.Section("recurrence").Key("frequency").Int()
	if err != nil {
		log.Fatalf("Failed to read parameter: %v", err)
	}
}

func OneInst(baseurl string) {

	price, err := priceFeed(symbol, baseurl, offset)
	errHandler(err)

	payload, err := payloadBuilder(symbol, amount, price, side)
	errHandler(err)

	signature := sigBuilder(payload)

	response, err := newOrder(baseurl, payload, signature)
	errHandler(err)

	log.Printf("%+v\n", response)
}

func MultiInst(baseurl string) {

	if freq <= 0 {
		log.Fatalln("enter frequency value greater than 0")

	} else {

		for {
			price, err := priceFeed(symbol, baseurl, offset)
			errHandler(err)

			payload, err := payloadBuilder(symbol, amount, price, side)
			errHandler(err)

			signature := sigBuilder(payload)

			response, err := newOrder(baseurl, payload, signature)
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

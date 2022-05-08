package main

import (
	"log"
	"time"
)

func oneInst(baseurl string, pretty bool) {

	price, err := PriceFeed(symbol, baseurl, offset)
	errHandler(err)

	payload, err := PayloadBuilder(symbol, amount, price, side)
	errHandler(err)

	signature := SigBuilder(payload, apisecret)

	response, err := NewOrder(baseurl, apikey, payload, signature)
	errHandler(err)

	if pretty {
		MakePretty(response)
	} else {
		log.Printf("%+v\n", response)
	}
}

func multiInst(baseurl string, freq int) {

	if freq <= 0 {
		log.Fatalln("enter frequency value greater than 0")

	} else {

		for {
			price, err := PriceFeed(symbol, baseurl, offset)
			errHandler(err)

			payload, err := PayloadBuilder(symbol, amount, price, side)
			errHandler(err)

			signature := SigBuilder(payload, apisecret)

			response, err := NewOrder(baseurl, apikey, payload, signature)
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

package main

import (
	"fmt"
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

	price, err := priceFeed(symbol, baseurl, offset)
	errHandler(err)

	payload, err := payloadBuilder(symbol, amount, price, side)
	errHandler(err)

	signature := sigBuilder(payload)

	response, err := newOrder(baseurl, payload, signature)
	errHandler(err)

	if pretty {
		fmt.Println(Pretty(response))
	} else {
		fmt.Printf("%+v\n", response)
	}
}

func MultiInst(baseurl string, freq int) {

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

func Pretty(r Response) string {
	resp := fmt.Sprintf(`
	OrderID:		%v
	ID:			%v
	Symbol:			%v
	Exchange:		%v
	AvgExecutionPrice:	%v
	Side:			%v
	Type:			%v
	Timestamp:		%v
	Timestampms:		%v
	IsLive:			%v
	IsCancelled:		%v
	IsHidden:		%v
	WasForced:		%v
	ExecutedAmount:		%v
	Options:		%v
	StopPrice:		%v
	Price:			%v
	OriginalAmount:		%v
	`, r.OrderID, r.ID, r.Symbol, r.Exchange,
		r.AvgExecutionPrice, r.Side, r.Type, r.Timestamp,
		r.Timestampms, r.IsLive, r.IsCancelled, r.IsHidden,
		r.WasForced, r.ExecutedAmount, r.Options, r.StopPrice,
		r.Price, r.OriginalAmount)

	return resp
}

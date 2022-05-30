package exec

import (
	"cgem/conf"
	"cgem/order"
	"fmt"
	"log"
	"os"
	"time"
)

func Execute(symbol, side string, amount float64, offset int) {

	c := conf.Get()

	switch c.Repeat {
	case true:
		if c.Freq <= 0 {
			logger(c.LogFile).Fatalln("enter frequency value greater than 0")
		} else {
			multiInst(symbol, side, amount, offset)
		}
	default:
		oneInst(symbol, side, amount, offset)
	}
}

func oneInst(symbol, side string, amount float64, offset int) {

	c := conf.Get()

	p, err := order.GetPrice(symbol, c.BaseURL)
	errHandler(c.LogFile, err)

	if c.Offset != 0 {
		offset = c.Offset
	}
	price, err := order.SetPrice(p.Price, offset)
	errHandler(c.LogFile, err)

	payload, err := order.PayloadBuilder(symbol, price, side, amount)
	errHandler(c.LogFile, err)

	signature := order.SigBuilder(payload, c.APISecret)

	response, err := order.New(c.BaseURL, payload, c.APIKey, signature)
	errHandler(c.LogFile, err)

	if c.Pretty {
		order.MakePretty(response)
		logger(c.LogFile).Printf("%+v\n", response)
	} else {
		fmt.Printf("%+v\n", response)
		logger(c.LogFile).Printf("%+v\n", response)
	}
}

func multiInst(symbol, side string, amount float64, offset int) {

	c := conf.Get()

	logger(c.LogFile).Println("app started")

	for {
		p, err := order.GetPrice(symbol, c.BaseURL)
		errHandler(c.LogFile, err)

		if c.Offset != 0 {
			offset = c.Offset
		}
		price, err := order.SetPrice(p.Price, offset)
		errHandler(c.LogFile, err)

		payload, err := order.PayloadBuilder(symbol, price, side, amount)
		errHandler(c.LogFile, err)

		signature := order.SigBuilder(payload, c.APISecret)

		response, err := order.New(c.BaseURL, payload, c.APIKey, signature)
		errHandler(c.LogFile, err)

		logger(c.LogFile).Printf("%+v\n", response)

		time.Sleep(time.Hour * time.Duration(c.Freq))
	}
}

func GetPrice(symbol string) {
	c := conf.Get()

	p, err := order.GetPrice(symbol, c.BaseURL)
	errHandler(c.LogFile, err)

	fmt.Printf("\n%v: %v\n", p.Pair, p.Price)
}

func logger(logfile string) *log.Logger {

	file := conf.GetPath(logfile)

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return log.New(f, "cgem: ", log.LstdFlags|log.Lshortfile)
}

func errHandler(logfile string, err error) {
	if err != nil {
		logger(logfile).Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/ini.v1"
)

const (
	PRODUCTION = "https://api.gemini.com"
	SANDBOX    = "https://api.sandbox.gemini.com"
)

var (
	amount    string
	apikey    string
	apisecret string
	env       string
	freq      int
	offset    int
	pretty    bool
	repeat    bool
	side      string
	symbol    string
)

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	env = cfg.Section("").Key("environment").String()
	apikey = cfg.Section("credentials").Key("apikey").String()
	apisecret = cfg.Section("credentials").Key("apisecret").String()

	pretty, err = cfg.Section("").Key("pretty").Bool()
	if err != nil {
		log.Fatalf("Failed to read parameter: %v", err)
	}
	repeat, err = cfg.Section("recurrence").Key("repeat").Bool()
	if err != nil {
		log.Fatalf("Failed to read parameter: %v", err)
	}
	freq, err = cfg.Section("recurrence").Key("frequency").Int()
	if err != nil {
		log.Fatalf("Failed to read parameter: %v", err)
	}
}

func oneInst(baseurl string) {

	g := GetPrice(symbol, baseurl)
	price, err := priceOffset(g, offset)
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

func multiInst(baseurl string) {

	if freq <= 0 {
		log.Fatalln("enter frequency value greater than 0")

	} else {

		for {
			g := GetPrice(symbol, baseurl)
			price, err := priceOffset(g, offset)
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

func getEnv(env string) string {
	switch env {
	case "production":
		return PRODUCTION
	case "sandbox":
		return SANDBOX
	default:
		return fmt.Sprintln(`enter a value of either "production" or "sandbox".`)
	}
}


func priceOffset (p *NewPrice, o int) (string, error) {
	price, err := strconv.ParseFloat(p.Price, 64)
		if err != nil {
			return "", fmt.Errorf("string convert parse float ecountered an error: %v", err)
	}
	return fmt.Sprint(price + float64(o)), nil
}
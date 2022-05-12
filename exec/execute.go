package exec

import (
	"fmt"
	"log"
	"os"
	"time"

	"cgem/order"
	"gopkg.in/ini.v1"
)

func init() {
	cfg, err := ini.Load(configFile)
	errHandler(err)

	env = cfg.Section("").Key("environment").String()
	apikey = cfg.Section("credentials").Key("apikey").String()
	apisecret = cfg.Section("credentials").Key("apisecret").String()
	logfile = cfg.Section("logging").Key("logfile").String()
	pretty, _ = cfg.Section("").Key("pretty").Bool()
	iOffset, _ = cfg.Section("orders").Key("offset").Int()
	repeat, _ = cfg.Section("recurrence").Key("repeat").Bool()
	freq, _ = cfg.Section("recurrence").Key("frequency").Int()
}

func Execute(symbol, side string, amount, offset int) {

	baseurl := getEnv(env)

	switch repeat {
	case true:
		if freq <= 0 {
			logger(logfile).Fatalln("enter frequency value greater than 0")
		} else {
			multiInst(symbol, side, baseurl, amount, offset)
		}
	default:
		oneInst(symbol, side, baseurl, amount, offset)
	}
}

func oneInst(symbol, side, baseurl string, amount, offset int) {

	p, err := order.PriceFeed(symbol, baseurl)
	errHandler(err)

	if iOffset != 0 {
		offset = iOffset
	}
	price, err := order.PriceOffset(p.Price, offset)
	errHandler(err)

	payload, err := order.PayloadBuilder(symbol, price, side, amount)
	errHandler(err)

	signature := order.SigBuilder(payload, apisecret)

	response, err := order.NewOrder(baseurl, apikey, payload, signature)
	errHandler(err)

	if pretty {
		order.MakePretty(response)
		logger(logfile).Printf("%+v\n", response)
	} else {
		fmt.Printf("%+v\n", response)
		logger(logfile).Printf("%+v\n", response)
	}
}

func multiInst(symbol, side, baseurl string, amount, offset int) {

	logger(logfile).Println("app started")

	for {
		p, err := order.PriceFeed(symbol, baseurl)
		errHandler(err)

		if iOffset != 0 {
			offset = iOffset
		}
		price, err := order.PriceOffset(p.Price, offset)
		errHandler(err)

		payload, err := order.PayloadBuilder(symbol, price, side, amount)
		errHandler(err)

		signature := order.SigBuilder(payload, apisecret)

		response, err := order.NewOrder(baseurl, apikey, payload, signature)
		errHandler(err)

		logger(logfile).Printf("%+v\n", response)

		time.Sleep(time.Hour * time.Duration(freq))
	}
}

func errHandler(err error) {
	if err != nil {
		logger(logfile).Fatalln(err)
	}
}

func getEnv(env string) string {
	switch env {
	case "production":
		return production
	default:
		return sandbox
	}
}

func GetPrice(symbol string) {
	baseurl := getEnv(env)

	p, err := order.PriceFeed(symbol, baseurl)
	errHandler(err)

	fmt.Printf("\n%v: %v\n", p.Pair, p.Price)
}

func logger(logfile string) *log.Logger {
	if logfile == "" {
		logfile = "cgem.log"
	}
	o, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return log.New(o, "cgem: ", log.LstdFlags|log.Lshortfile)
}

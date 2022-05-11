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
	cfg, err := ini.Load("config.ini")
	errHandler(err)

	env = cfg.Section("").Key("environment").String()
	apikey = cfg.Section("credentials").Key("apikey").String()
	apisecret = cfg.Section("credentials").Key("apisecret").String()
	pretty, _ = cfg.Section("").Key("pretty").Bool()
	iniOffset, _ = cfg.Section("orders").Key("offset").Int()
	repeat, _ = cfg.Section("recurrence").Key("repeat").Bool()
	freq, _ = cfg.Section("recurrence").Key("frequency").Int()

}


func Execute(symbol, side string, amount, offset int) {

	if iniOffset != 0 {
		offset = iniOffset
	}

	switch repeat {
	case true:
		multiInst(symbol, side, amount, offset)
	default:
		oneInst(symbol, side,amount, offset)
	}
}



func oneInst(symbol, side string, amount, offset int) {

	baseurl := getEnv(env)

	p, err := order.PriceFeed(symbol, baseurl)
	errHandler(err)

	price, err := order.PriceOffset(p.Price, offset)
	errHandler(err)

	payload, err := order.PayloadBuilder(symbol, price, side, amount)
	errHandler(err)

	signature := order.SigBuilder(payload, apisecret)

	response, err := order.NewOrder(baseurl, apikey, payload, signature)
	errHandler(err)

	if pretty {
		order.MakePretty(response)
		logger().Printf("%+v\n", response)
	} else {
		fmt.Printf("%+v\n", response)
		logger().Printf("%+v\n", response)
	}
}

func multiInst(symbol, side string, amount, offset int) {

	baseurl := getEnv(env)

	if freq <= 0 {
		logger().Fatalln("enter frequency value greater than 0")

	} else {

		logger().Println("app started")

		for {

			p, err := order.PriceFeed(symbol, baseurl)
			errHandler(err)

			price, err := order.PriceOffset(p.Price, offset)
			errHandler(err)

			payload, err := order.PayloadBuilder(symbol, price, side, amount)
			errHandler(err)

			signature := order.SigBuilder(payload, apisecret)

			response, err := order.NewOrder(baseurl, apikey, payload, signature)
			errHandler(err)

			logger().Printf("%+v\n\n", response)

			time.Sleep(time.Hour * time.Duration(freq))
		}
	}
}

func errHandler(err error) {
	if err != nil {
		logger().Fatalln(err)
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

	fmt.Printf("%v: %v\n", p.Pair, p.Price)
}

func logger() *log.Logger {
	logfile, err := os.OpenFile("cgem.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return log.New(logfile, "cgem: ", log.LstdFlags|log.Lshortfile)
}
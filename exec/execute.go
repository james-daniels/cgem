package exec

import (
	"log"
	"time"

	"cgem/order"
	"gopkg.in/ini.v1"
)

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	env = cfg.Section("").Key("environment").String()
	apikey = cfg.Section("credentials").Key("apikey").String()
	apisecret = cfg.Section("credentials").Key("apisecret").String()

	// Optional parameters
	pretty, _ = cfg.Section("").Key("pretty").Bool()
	iniOffset, _ = cfg.Section("orders").Key("offset").Int()
	repeat, _ = cfg.Section("recurrence").Key("repeat").Bool()
	freq, _ = cfg.Section("recurrence").Key("frequency").Int()

}


func Execute(symbol, amount, side string, offset int) {

	if iniOffset != 0 {
		offset = iniOffset
	}

	switch repeat {
	case true:
		multiInst(symbol, amount, side, offset)
	default:
		oneInst(symbol, amount, side, offset)
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

func oneInst(symbol, amount, side string, offset int) {

	baseurl := getEnv(env)

	gp := order.GetPrice(symbol, baseurl)
	price, err := order.PriceOffset(gp.Price, offset)
	errHandler(err)

	payload, err := order.PayloadBuilder(symbol, amount, price, side)
	errHandler(err)

	signature := order.SigBuilder(payload, apisecret)

	response, err := order.NewOrder(baseurl, apikey, payload, signature)
	errHandler(err)

	if pretty {
		order.MakePretty(response)
	} else {
		log.Printf("%+v\n", response)
	}
}

func multiInst(symbol, amount, side string, offset int) {

	baseurl := getEnv(env)

	if freq <= 0 {
		log.Fatalln("enter frequency value greater than 0")

	} else {

		for {
			gp := order.GetPrice(symbol, baseurl)
			price, err := order.PriceOffset(gp.Price, offset)
			errHandler(err)

			payload, err := order.PayloadBuilder(symbol, amount, price, side)
			errHandler(err)

			signature := order.SigBuilder(payload, apisecret)

			response, err := order.NewOrder(baseurl, apikey, payload, signature)
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
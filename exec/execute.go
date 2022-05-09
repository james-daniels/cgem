package exec

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/ini.v1"
	"cgem/control"
)

const (
	production = "https://api.gemini.com"
	sandbox    = "https://api.sandbox.gemini.com"
)

var (
	apikey    string
	apisecret string
	env       string
	freq      int
	pretty    bool
	repeat    bool
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


func Execute(symbol, amount, side string, offset int) {
	
	switch repeat {
	case true:
		multiInst(symbol, amount, side, offset)
	default:
		oneInst(symbol, amount, side, offset)
	}
}


func oneInst(symbol, amount, side string, offset int) {

	baseurl := getEnv(env)

	gp := control.GetPrice(symbol, baseurl)
	price, err := priceOffset(gp.Price, offset)
	errHandler(err)

	payload, err := control.PayloadBuilder(symbol, amount, price, side)
	errHandler(err)

	signature := control.SigBuilder(payload, apisecret)

	response, err := control.NewOrder(baseurl, apikey, payload, signature)
	errHandler(err)

	if pretty {
		control.MakePretty(response)
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
			gp := control.GetPrice(symbol, baseurl)
			price, err := priceOffset(gp.Price, offset)
			errHandler(err)

			payload, err := control.PayloadBuilder(symbol, amount, price, side)
			errHandler(err)

			signature := control.SigBuilder(payload, apisecret)

			response, err := control.NewOrder(baseurl, apikey, payload, signature)
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
		return production
	case "sandbox":
		return sandbox
	default:
		return fmt.Sprintln(`enter a value of either "production" or "sandbox".`)
	}
}

func priceOffset(price string, offset int) (string, error) {
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return "", fmt.Errorf("string convert parse float ecountered an error: %v", err)
	}
	return fmt.Sprint(p + float64(offset)), nil
}
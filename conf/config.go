package conf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/ini.v1"
)

const (
	configFile = "config.ini"
)

type config struct {
	Env       string
	APIKey    string
	APISecret string
	Freq      int
	Offset    int
	LogFile   string
	Pretty    bool
	Repeat    bool
	BaseURL   string
}

func Builder() *config {
	return &config{}
}

func (c *config) Set(env, apiKey, apiSecret string) {
	c.Env = env
	c.APIKey = apiKey
	c.APISecret = apiSecret
}

func Build(conf *config) {

	configTemplate := `
#Possible values: sandbox and production
environment = {{.Env}}

#Optional: Present output in human readable format
#Only available for single run jobs
pretty = true

[credentials]
#API key and secret
apikey = {{.APIKey}}
apisecret = {{.APISecret}}

[recurrence]
#Optional: Only for recurring jobs
repeat = false

#Dependent on repeat = true
#Number of hours between runs
frequency = 0

[orders]
#Default value is 0
#The API does not support market orders because it does not provide price protection.
#Offset agressively coupled with the curret price increases or decreases the limit price.
#This will achieve the same result as a market order.
#offset = 0

[logging]
#Optional: path to log file location
#logfile = "cgem.log"
`

	f, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	t := template.Must(template.New("configTemplate").Parse(configTemplate))
	err = t.Execute(f, conf)
	if err != nil {
		log.Println("an error has occured with config template")
	}

	fAbs, err := filepath.Abs(f.Name())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("created config file:", fAbs)
}

func Get() *config {

	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln(configFile, "missing: run 'cgem init' to get started")
		}
	}

	var (
		apiKey    string
		apiSecret string
		env       string
		freq      int
		offset    int
		logFile   string
		pretty    bool
		repeat    bool
	)

	cfg, err := ini.Load(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	apiKey = cfg.Section("credentials").Key("apikey").String()
	apiSecret = cfg.Section("credentials").Key("apisecret").String()
	env = cfg.Section("").Key("environment").String()
	logFile = cfg.Section("logging").Key("logfile").String()
	if logFile == "" {
		logFile = "cgem.log"
	}
	pretty, _ = cfg.Section("").Key("pretty").Bool()
	offset, _ = cfg.Section("orders").Key("offset").Int()
	repeat, _ = cfg.Section("recurrence").Key("repeat").Bool()
	freq, _ = cfg.Section("recurrence").Key("frequency").Int()

	return &config{
		Env:       env,
		APIKey:    apiKey,
		APISecret: apiSecret,
		Freq:      freq,
		Offset:    offset,
		LogFile:   logFile,
		Pretty:    pretty,
		Repeat:    repeat,
		BaseURL:   getEnv(env),
	}
}

func getEnv(env string) string {

	switch env {
	case "production":
		return "https://api.gemini.com"
	case "sandbox":
		return "https://api.sandbox.gemini.com"
	default:
		return "enter a valid environment: sandbox or production"
	}
}
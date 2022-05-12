package exec

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type configBuilder struct {
	Env       string
	APIKey    string
	APISecret string
}

func NewConfigBuilder() *configBuilder {
	return &configBuilder{}
}

func (c *configBuilder) SetConfig(env, apiKey, apiSecret string) {
	c.Env = env
	c.APIKey = apiKey
	c.APISecret = apiSecret
}

func (c configBuilder) BuildConfig(conf *configBuilder) {

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

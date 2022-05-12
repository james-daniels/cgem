package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init creates the default config file",
	Long:  "Init creates the default config file in the current binary location",
	Run: func(cmd *cobra.Command, args []string) {

		conf := newConfigBuilder()
		conf.setConfig(env, apiKey, apiSecret)
		buildConfig(conf)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&env, "env", "e", "", "enter the environment values: sandbox or production")
	initCmd.MarkFlagRequired("env")
	initCmd.Flags().StringVarP(&apiKey, "key", "k", "", "enter the api key")
	initCmd.MarkFlagRequired("key")
	initCmd.Flags().StringVarP(&apiSecret, "secret", "s", "", "enter the api secret")
	initCmd.MarkFlagRequired("secret")
}

type configBuilder struct {
	Environment string
	APIKey string
	APISecret string
}

func newConfigBuilder() *configBuilder {
	return &configBuilder{}
}

func (c *configBuilder) setConfig(env, apiKey, apiSecret string) {
	c.Environment = env
	c.APIKey = apiKey
	c.APISecret = apiSecret
}

func buildConfig(c *configBuilder) {

	configTemplate :=`
#Possible values: sandbox and production
environment = {{.Environment}}

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
	err = t.Execute(f, c)
	if err != nil {
		log.Println("an error has occured with config template")
	}

	fAbs, err := filepath.Abs(f.Name())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("created config file:", fAbs)
}
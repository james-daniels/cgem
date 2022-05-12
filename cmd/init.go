package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const configFile = "config.ini"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init creates the default config file",
	Long:  "Init creates the default config file in the current binary location",
	Run: func(cmd *cobra.Command, args []string) {

		loadConfigFile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func loadConfigFile() {

	configTemplate := `
#Possible values: sandbox and production
environment = sandbox

#Optional: Present output in human readable format
#Only available for single run jobs
#pretty = true

[credentials]
#API key and secret
apikey = account-XXXXXXXXXXXXXXXXXXXX
apisecret = XXXXXXXXXXXXXXXXXXXX

[recurrence]
#Optional: Only for recurring jobs
#repeat = false

#Dependent on repeat = true
#Number of hours between runs
#frequency = 0

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
	file, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	file.Write([]byte(configTemplate))

	fAbs, err := filepath.Abs(file.Name())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("created config file:", fAbs)
}

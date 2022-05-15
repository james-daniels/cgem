package cmd

import (
	"cgem/conf"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init creates the default config file",
	Long:  "Init creates the default config file in the current binary location",
	Run: func(cmd *cobra.Command, args []string) {

		config := conf.Builder()
		config.Set(env, apiKey, apiSecret)
		conf.Build(config)
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

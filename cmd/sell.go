/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"cgem/exec"
)

// sellCmd represents the sell command
var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "A brief description of your command",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {

	exec.Execute(symbol, amount, sside, offset)

	},
}

func init() {
	rootCmd.AddCommand(sellCmd)

	sellCmd.PersistentFlags().StringVarP(&symbol, "symbol","s", "", "SYMBOL: symbol for the new order")
	sellCmd.PersistentFlags().StringVarP(&amount, "amount","a", "", "AMOUNT: amount to purchase")
	sellCmd.PersistentFlags().IntVarP(&offset, "offset","o", 0, "OFFSET: amount to ADD TO PRICE")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"cgem/exec"
)

// buyCmd represents the buy command
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Short description to buy",
	Long: "Long description to use the buy function on the app",
	Run: func(cmd *cobra.Command, args []string) {

	exec.Execute(symbol, amount, bside, offset)

	},
}

func init() {
	rootCmd.AddCommand(buyCmd)

	buyCmd.Flags().StringVarP(&symbol, "symbol","s", "", "SYMBOL: symbol for the new order")
	buyCmd.Flags().StringVarP(&amount, "amount","a", "", "AMOUNT: amount to purchase")
	buyCmd.Flags().IntVarP(&offset, "offset","o", 0, "OFFSET: amount to ADD TO PRICE")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cgem/exec"
	"github.com/spf13/cobra"
)

// buyCmd represents the buy command
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "buy side places order to buy crypto",
	Long: "buy will fill whatever part of the order it can immediately, then cancel any remaining amount",
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

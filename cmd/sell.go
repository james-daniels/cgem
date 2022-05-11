/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cgem/exec"
	"github.com/spf13/cobra"
)

// sellCmd represents the sell command
var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "sell side places order to buy crypto",
	Long: "sell will fill whatever part of the order it can immediately, then cancel any remaining amount",
	Run: func(cmd *cobra.Command, args []string) {

	exec.Execute(symbol, amount, sside, offset)

	},
}

func init() {
	rootCmd.AddCommand(sellCmd)

	sellCmd.Flags().StringVarP(&symbol, "symbol","s", "", "SYMBOL: symbol for the new order")
	sellCmd.Flags().StringVarP(&amount, "amount","a", "", "AMOUNT: amount to purchase")
	sellCmd.Flags().IntVarP(&offset, "offset","o", 0, "OFFSET: amount to ADD TO PRICE")
}

package cmd

import (
	"cgem/exec"

	"github.com/spf13/cobra"
)

var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Price of trading pair",
	Long: "Get the price of trading pair by providing the symbol",
	Run: func(cmd *cobra.Command, args []string) {

		exec.GetPrice(symbol)

	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	priceCmd.Flags().StringVarP(&symbol, "symbol","s", "", "SYMBOL: symbol of the trading pair")
	priceCmd.MarkFlagRequired("symbol")
}

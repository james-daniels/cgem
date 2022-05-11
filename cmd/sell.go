package cmd

import (
	"cgem/exec"

	"github.com/spf13/cobra"
)

var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "Sell side places order to buy crypto",
	Long: "Sell will fill whatever part of the order it can immediately, then cancel any remaining amount",
	Run: func(cmd *cobra.Command, args []string) {

		exec.Execute(symbol, sside, amount, offset)
	},
}

func init() {
	rootCmd.AddCommand(sellCmd)

	sellCmd.Flags().StringVarP(&symbol, "symbol","s", "", "SYMBOL: symbol of the trading pair")
	sellCmd.MarkFlagRequired("symbol")
	sellCmd.Flags().IntVarP(&amount, "amount","a", 0, "AMOUNT: amount to purchase")
	sellCmd.MarkFlagRequired("amount")
	sellCmd.Flags().IntVarP(&offset, "offset","o", 0, "OFFSET: amount to ADD TO PRICE")
}

package cmd

import (
	"cgem/exec"

	"github.com/spf13/cobra"
)

var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Buy side places order to buy crypto",
	Long:  "Buy will fill part of the order it can immediately, then cancel any remaining amount.",
	Run: func(cmd *cobra.Command, args []string) {

		exec.Execute(symbol, bside, amount, offset)
	},
}

func init() {
	rootCmd.AddCommand(buyCmd)

	buyCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "symbol of the trading pair")
	buyCmd.MarkFlagRequired("symbol")
	buyCmd.Flags().IntVarP(&amount, "amount", "a", 0, "amount to buy")
	buyCmd.MarkFlagRequired("amount")
	buyCmd.Flags().IntVarP(&offset, "offset", "o", 0, "positive value to add to price")
}

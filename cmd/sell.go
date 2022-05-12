package cmd

import (
	"cgem/exec"

	"github.com/spf13/cobra"
)

var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "Sell side places order to sell crypto",
	Long:  "Sell will fill part of the order it can immediately, then cancel any remaining amount.",
	Run: func(cmd *cobra.Command, args []string) {

		exec.Execute(symbol, sside, amount, offset)
	},
}

func init() {
	rootCmd.AddCommand(sellCmd)

	sellCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "symbol of the trading pair")
	sellCmd.MarkFlagRequired("symbol")
	sellCmd.Flags().IntVarP(&amount, "amount", "a", 0, "amount to sell")
	sellCmd.MarkFlagRequired("amount")
	sellCmd.Flags().IntVarP(&offset, "offset", "o", 0, "amount to add")
}

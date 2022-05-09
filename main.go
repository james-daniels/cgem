package main

import (
	"flag"
)

func init() {
	flag.StringVar(&symbol, "s", "", "SYMBOL: symbol for the new order")
	flag.StringVar(&symbol, "symbol", "", "SYMBOL: symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: amount to purchase")
	flag.StringVar(&amount, "amount", "", "AMOUNT: amount to purchase")
	flag.IntVar(&offset, "o", 0, `OFFSET: amount to ADD TO PRICE (default "0")`)
	flag.IntVar(&offset, "offset", 0, `OFFSET: amount to ADD TO PRICE (default "0")`)
	flag.StringVar(&side, "S", "buy", "SIDE TYPE: buy or sell")
	flag.StringVar(&side, "side", "buy", "SIDE TYPE: buy or sell")
}

func main() {
	flag.Parse()

	switch repeat {
	case true:
		multiInst()
	default:
		oneInst()
	}
}

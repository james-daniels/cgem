package main

import (
	"fmt"
)

const (
	PRODURL = "https://api.gemini.com"
	SANDURL = "https://api.sandbox.gemini.com"
)

func main() {

	payload, _ := PayloadBuilder("ltcusd", "1", "105.00", "buy")

	signature := SignatureBuilder(payload)

	result, _ := PostOrder(SANDURL, payload, signature)

	fmt.Printf("%+v\n", result)
}

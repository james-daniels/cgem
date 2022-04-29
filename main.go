package main

import (
	"fmt"
)

func main() {

	payload, _ := PayloadBuilder("ltcusd", "1", "105.00")

	signature := SignatureBuilder(payload)

	result, _ := PostOrder(SANDURL, payload, signature)

	fmt.Printf("%+v\n", result)

}

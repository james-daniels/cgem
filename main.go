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

//sources:
//https://golangbyexample.com/set-headers-http-request/
//https://golangdocs.com/json-with-golang
//https://golangbyexample.com/base64-golang/
//https://golangcode.com/generate-sha256-hmac/
//https://docs.gemini.com/rest-api/#new-order

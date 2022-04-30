package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	APIKEY    = os.Getenv("apikey")
	APISECRET = os.Getenv("apisecret")
)

type NewPayload struct {
	Request string   `json:"request"`
	Nonce   string   `json:"nonce"`
	Symbol  string   `json:"symbol"`
	Amount  string   `json:"amount"`
	Price   string   `json:"price"`
	Side    string   `json:"side"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

type Response struct {
	OrderID           string   `json:"order_id"`
	ID                string   `json:"id"`
	Symbol            string   `json:"symbol"`
	Exchange          string   `json:"exchange"`
	AvgExecutionPrice string   `json:"avg_execution_price"`
	Side              string   `json:"side"`
	Type              string   `json:"type"`
	Timestamp         string   `json:"timestamp"`
	Timestampms       int      `json:"timestampms"`
	IsLive            bool     `json:"is_live"`
	IsCancelled       bool     `json:"is_cancelled"`
	IsHidden          bool     `json:"is_hidden"`
	WasForced         bool     `json:"was_forced"`
	ExecutedAmount    string   `json:"executed_amount"`
	Options           []string `json:"options"`
	StopPrice         string   `json:"stop_price"`
	Price             string   `json:"price"`
	OriginalAmount    string   `json:"original_amount"`
}

func PayloadBuilder(symbol, amount, price, side string) (string, error) {

	nonce := fmt.Sprint(time.Now().Unix() * 1000)

	p := &NewPayload{
		Request: "/v1/order/new",
		Nonce:   nonce,
		Symbol:  symbol,
		Amount:  amount,
		Price:   price,
		Side:    side,
		Type:    "exchange limit",
		Options: []string{
			"immediate-or-cancel",
		},
	}

	encodePayload, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("ecountered an error: %v", err)
	}
	payload := base64.StdEncoding.EncodeToString(encodePayload)

	return payload, nil
}

func SigBuilder(payload string) string {

	h := hmac.New(sha512.New384, []byte(APISECRET))
	h.Write([]byte(payload))

	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}

func PostOrder(baseurl, payload, signature string) (Response, error) {

	endpoint := "/v1/order/new"
	url := baseurl + endpoint

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return Response{}, fmt.Errorf("ecountered an error: %v", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Add("Content-Length", "0")
	req.Header.Add("X-GEMINI-APIKEY", APIKEY)
	req.Header.Add("X-GEMINI-PAYLOAD", payload)
	req.Header.Add("X-GEMINI-SIGNATURE", signature)
	req.Header.Add("Cache-Control", "no-cache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("ecountered an error: %v", err)
	}
	defer resp.Body.Close()

	var response Response
	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return Response{}, fmt.Errorf("ecountered an error: %v", err)
		}
	} else {
		resp.Body.Close()
		return Response{}, fmt.Errorf("%v: ecountered an error: %v", resp.StatusCode, err)
	}

	return response, nil
}

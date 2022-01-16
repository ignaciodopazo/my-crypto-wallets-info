package binanceEndpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ErrorReceived struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"msg"`
}

func Account(APIKey, secretKey string) []byte {

	b := NewURLBuilder()
	b.SetBaseURL("https://api.binance.com/api/v3/account")
	b.AddQueryParam("timestamp", fmt.Sprint(time.Now().UnixMilli()))
	b.addSignature(secretKey)

	accountInfoUrl := b.GenerateURL()

	req, err := http.NewRequest("GET", accountInfoUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Mbx-Apikey", APIKey)

	client := &http.Client{Timeout: time.Minute}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var errBody ErrorReceived
	// is not in statusCode, but in body

	err = json.Unmarshal(body, &errBody)
	if err != nil {
		log.Fatal(err)
	}

	// msg: "Timestamp for this request is outside of the recvWindow."
	if errBody.ErrorCode == -1021 {
		Account(APIKey, secretKey)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	return body
}

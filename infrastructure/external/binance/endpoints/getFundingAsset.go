package binanceEndpoint

import (
	_ "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetFundingAsset(APIKey, secretKey string) []byte {

	b := NewURLBuilder()
	b.SetBaseURL("https://api.binance.com/sapi/v1/asset/get-funding-asset")
	b.AddQueryParam("timestamp", fmt.Sprint(time.Now().UnixMilli()))
	b.addSignature(secretKey)

	getFundingAssetUrl := b.GenerateURL()

	req, err := http.NewRequest("POST", getFundingAssetUrl, nil)
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

	if resp.StatusCode != 200 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	return body
}

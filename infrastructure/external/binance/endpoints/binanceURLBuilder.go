package binanceEndpoint

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"

	dom "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

// Make it implement URLBuilder interface
type binanceURLBuilder struct {
	baseURL string
    queryParams dom.QueryParams
}

func NewURLBuilder() *binanceURLBuilder {
	return &binanceURLBuilder{
		baseURL:     "",
		queryParams: dom.NewQueryParams(),
	}
}

func (b *binanceURLBuilder) SetBaseURL(url string)  {
	b.baseURL = url
}

func (b *binanceURLBuilder) AddQueryParam(key, value string)  {
	b.queryParams.AddQueryParam(key, value)
}

// Implementation dependent of Binance API constraint about position of signature in query params:
// signature must ALWAYS be as the last query parameter.
func (b *binanceURLBuilder) GenerateURL() string  {
	if b.queryParams == nil || b.baseURL == "" {
		return b.baseURL
	}
	result := b.baseURL
	firstPass := true
	hasSignature := false
	for k, v := range b.queryParams {
		if k == "signature" {
			hasSignature = true
			continue
		}
		if !firstPass {
			result = result + "&" + k + "=" + v
		} else {
			result = result + "?" + k + "=" + v
			firstPass = false
		}
	}
	if hasSignature {
		result = result + "&" + "signature" + "=" + b.queryParams["signature"]
	}
	return result
}

// Now binance URLBuilder implements URLBuilder interface

// Utility for addSignature
func (b *binanceURLBuilder) getTotalParams() string {
	result := ""
	firstPass := true
	for k, v := range b.queryParams {
		if !firstPass {
			result = result + "&" + k + "=" + v
		} else {
			result = k + "=" + v
			firstPass = false
		}
	}
	return result
}

func (b *binanceURLBuilder) addSignature(secretKey string)  {
	totalParams := b.getTotalParams()
	signature, err := newSignature(totalParams, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	b.queryParams.AddQueryParam("signature", signature)
}

// For SIGNED endpoints, as specified in
// https://binance-docs.github.io/apidocs/spot/en/#signed-trade-user_data-and-margin-endpoint-security.
func newSignature(totalParams, secretKey string) (string, error) {
	h := hmac.New(sha256.New, []byte(secretKey))
	_, err := h.Write([]byte(totalParams))
	if err != nil {
		return "", err
	}
	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil
}

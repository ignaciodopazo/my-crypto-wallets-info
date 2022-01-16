package geckoController

import (
	e "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

type CoinInfo interface {
	GetName() string
	GetCurrentPrice() float64
}

func GetCoinsPrices(coinList []CoinInfo) e.AssetPrices {

	result := e.NewAssetPrice()
	for _, coin := range coinList {
		result.AddAssetPrice(coin.GetName(), coin.GetCurrentPrice())
	}

	return result
}

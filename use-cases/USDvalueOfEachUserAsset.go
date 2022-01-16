package usecase

import (
	e "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

// Returs the value in USD for each user's asset.
func GetUSDValueOfEachUserAsset(userAssets e.UserAssets, assetPrices e.AssetPrices) map[string]float64 {
	return e.USDValueOfEach(userAssets, assetPrices)
}

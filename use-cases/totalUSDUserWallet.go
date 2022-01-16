package usecase

import (
	e "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

// Returns the total funds of user wallet
func GetTotalUSDUserWallet(userAssets e.UserAssets, assetPrices e.AssetPrices) float64 {
	return e.UserAssetsTotalValue(userAssets, assetPrices)
}

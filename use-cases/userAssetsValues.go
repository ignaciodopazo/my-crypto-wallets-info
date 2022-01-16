package usecase

import (
	e "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

// Returns each user asset value.
func GetUserAssetsValues(userAssets e.UserAssets) map[string]float64 {
	return userAssets
}

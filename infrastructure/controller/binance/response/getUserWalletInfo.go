package binanceController

import (
	"encoding/json"
	"log"

	f "github.com/ignaciodopazo/funTools/src/v0"

	. "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
	"github.com/ignaciodopazo/my-crypto-wallets-info/utils/srcutils"
)

// Type to parse the response of get-funding-asset binance endpoint
type AssetBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// Type to parse the response of account binance endpoint
type AccountInfo struct {
	AccountType string         `json:"accountType"`
	Balances    []AssetBalance `json:"balances"`
}

// Receives the responses of binance's get-funding-asset and accound endpoints.
//
// Returns the user assets with its corresponding amount.
func GetUserWalletInfo(accountResponse, getFundingAssetResponse []byte) UserAssets {

	var accInfo AccountInfo
	// gives the assets quantity from SPOT and EARN
	err := json.Unmarshal(accountResponse, &accInfo)
	if err != nil {
		log.Fatal(err)
	}

	// filter the currencies of which the user doesn't possess any amount
	balances := f.Filter(func(ab AssetBalance) bool {
		free := srcutils.ParseFloat(ab.Free)
		locked := srcutils.ParseFloat(ab.Locked)
		return free != 0 || locked != 0
	}, accInfo.Balances)

	// Transform those that are in EARN (LD prefix in asset's name).
	// If there exists an asset that begins with LD, it breaks
	accInfo.Balances = f.Fmap(func(ab AssetBalance) AssetBalance {
		if ab.Asset[:2] == "LD" {
			ab.Asset = ab.Asset[2:]
		}
		return ab
	}, balances)

	var fundingAssetBalances []AssetBalance
	// gives the assets quantity from FUNDING
	err = json.Unmarshal(getFundingAssetResponse, &fundingAssetBalances)
	if err != nil {
		log.Fatal(err)
	}

	// join balance of EARN and SPOT with FUNDING
	accountTotalBalance := append(accInfo.Balances, fundingAssetBalances...)

	userAssets := NewUserAssets()
	for _, ab := range accountTotalBalance {
		userAssets.AddAsset(ab.Asset, srcutils.ParseFloat(ab.Free)+srcutils.ParseFloat(ab.Locked))
	}

	return userAssets
}

package main

import (
	"fmt"

	app "github.com/ignaciodopazo/my-crypto-wallets-info/app"

	be "github.com/ignaciodopazo/my-crypto-wallets-info/infrastructure/external/binance/endpoints"
	ge "github.com/ignaciodopazo/my-crypto-wallets-info/infrastructure/external/gecko/endpoints"

	bc "github.com/ignaciodopazo/my-crypto-wallets-info/infrastructure/controller/binance/response"
	gc "github.com/ignaciodopazo/my-crypto-wallets-info/infrastructure/controller/gecko/response"

	"github.com/ignaciodopazo/my-crypto-wallets-info/use-cases"

	p "github.com/ignaciodopazo/my-crypto-wallets-info/infrastructure/presenter"
)

func main() {
	// configuration checking, cli
	userConfig := app.GetUserConfig()
	buc := userConfig.GetBinanceAccountConfig()

	// binance API endpoints
	accountRes := be.Account(buc.APIKey, buc.SecretKey)
	getFundingAssetRes := be.GetFundingAsset(buc.APIKey, buc.SecretKey)
	// gecko API endpoint
	coinsMarketRes := ge.CoinsMarket()

	// binance API controller
	userAssets := bc.GetUserWalletInfo(accountRes, getFundingAssetRes)
	// gecko API controller
	assetsPrices := gc.GetCoinsPrices(coinsMarketRes)

	// use cases
	userAssetValues := usecase.GetUserAssetsValues(userAssets)
	userUSDValPerAsset := usecase.GetUSDValueOfEachUserAsset(userAssets, assetsPrices)
	userTotalUSD := usecase.GetTotalUSDUserWallet(userAssets, assetsPrices)

	// presenters
	toPrint1 := p.ShowUserWalletAssetsToUSDValue(userAssetValues, userUSDValPerAsset)
	toPrint2 := p.ShowUserWalletTotalUSDValue(userTotalUSD)

	fmt.Println(toPrint1)
	fmt.Println()
	fmt.Println(toPrint2)
	fmt.Println()
}

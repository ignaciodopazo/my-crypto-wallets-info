package geckoEndpoint

import (
	"log"
	"strings"

	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"

	gc "github.com/ignaciodopazo/my-crypto-wallets-info/infrastructure/controller/gecko/response"
)

// Make it implement CoinInfo interface
type geckoCoinInfo struct {
	name  string
	price float64
}

func (gci geckoCoinInfo) GetName() string {
	return gci.name
}

func (gci geckoCoinInfo) GetCurrentPrice() float64 {
	return gci.price
}

// From now on, geckoCoinInfo

// Main function of the module
func CoinsMarket() []gc.CoinInfo {
	cg := gecko.NewClient(nil)

	var (
		allCoins       []string = nil
		coinsInOnePage int      = 2000
	)
	priceChangePercentageAt := []string{"1h", "24h", "7d", "14d", "30d", "200d", "1y"}

	cm, err := cg.CoinsMarket("usd", allCoins, "market_cap_desc", coinsInOnePage,
		1, false, priceChangePercentageAt)
	if err != nil {
		log.Fatal(err)
	}
	return toCoinInfo(cm)
}

func toCoinInfo(cm *types.CoinsMarket) []gc.CoinInfo {
	coinList := *cm
	result := []gc.CoinInfo{}
	for _, coin := range coinList {
		result = append(result, geckoCoinInfo{name: strings.ToUpper(coin.Symbol),
			price: coin.CurrentPrice})
	}
	return result
}

package presenter

import (
	"fmt"
)

type UserAssetsValues map[string]float64

type UserAssetsValueInUSD map[string]float64

// Receives as arguments the result of GetUserAssetsValues and GetUSDValueOfEachUserAsset use cases
// in that order.
//
// Returns a log to be printed in CLI.
func ShowUserWalletAssetsToUSDValue(uav UserAssetsValues, uavInUSD UserAssetsValueInUSD) (log string) {
	for name, coinAmount := range uav {
		usdVal := uavInUSD[name]
		if usdVal == 0 {
			fmt.Println("\n", name, "NOT FOUND IN COIN LIST")
			continue
		}
		log += fmt.Sprint("\n ",
			fmt.Sprintf("%f", coinAmount),
				" " + name +
				" at " + fmt.Sprintf("%.2f", usdVal/coinAmount) +
				" USD/" + name + " -> ",
				fmt.Sprintf("%.2f", usdVal), " USD")
	}
	return log
}

// Receives as argument the result of GetTotalUSDUserWallet use case.
//
// Returns a log to be printed in CLI.
func ShowUserWalletTotalUSDValue(totalValue float64) (log string) {
	return "Current total value: " + fmt.Sprintf("%.2f", totalValue) + " USD"
}

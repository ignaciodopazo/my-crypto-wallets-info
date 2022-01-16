package domain


// Returns a map with each user's asset asociated with its value in USD.
func USDValueOfEach(ua UserAssets, ap AssetPrices) map[string]float64 {
	result := map[string]float64{}
	for name, quantity := range ua {
		result[name] = quantity * ap[name]
	}
	return result
}

// Returns the user asset's total value in USD
func UserAssetsTotalValue(ua UserAssets, ap AssetPrices) float64 {
	var result float64 = 0
	for name, quantity := range ua {
		result += quantity * ap[name]
	}
	return result
}

// Returns both results of USDValueOfEach and UserAssetsTotalValue
// packed in a tuple (in that order). For efficiency reasons you should
// use this function if you need both.
func (ua UserAssets) UserAssetsValues(ap AssetPrices) (map[string]float64, float64) {
	var (
		USDValuesOfEach         = map[string]float64{}
		totalValue      float64 = 0
	)
	for name, quantity := range ua {
		totalValueOfAsset := quantity * ap[name]
		USDValuesOfEach[name] = totalValueOfAsset
		totalValue += totalValueOfAsset
	}
	return USDValuesOfEach, totalValue
}
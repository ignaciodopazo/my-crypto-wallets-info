package domain

// Type to represent the current value of assets.
type AssetPrices map[string]float64

func NewAssetPrice() AssetPrices {
	return map[string]float64{}
}

// Add an asset and its corresponding price to the structure. If the
// pair already exists, it updates the price.
func (ap AssetPrices) AddAssetPrice(name string, quantity float64) {
	ap[name] = quantity
}

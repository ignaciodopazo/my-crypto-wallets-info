package domain

// An user asset has a name associated with a quantity.
// Follow the convention to put asset's name in upper case.
type UserAssets map[string]float64

func NewUserAssets() UserAssets {
	return map[string]float64{}
}

// If the given asset already exists in the structure, it adds
// to it the quantity given as argument. Otherwise, it adds the
// new asset to the structure.
func (ua UserAssets) AddAsset(name string, quantity float64) {
	currentQuantity := ua[name]
	if currentQuantity != 0 {
		ua[name] = currentQuantity + quantity
	} else {
		ua[name] = quantity
	}
}

// Returns the quantity of the given asset.
func (ua UserAssets) AssetQuantity(name string) (quantity float64) {
	quantity = ua[name]
	return
}

// Deletes the asset represented by name, returning its associated value.
func (ua UserAssets) PopAsset(name string) (quantity float64) {
	quantity, ok := ua[name]
	if ok {
		delete(ua, name)
	}
	return
}

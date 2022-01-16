package test

import (
	"reflect"
	"testing"

	. "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)


func TestUSDValues(t *testing.T) {

	t.Run("USDValueOfEach", func(t *testing.T) {
		ua := NewUserAssets()
		ua.AddAsset("A", 10.0)
		ua.AddAsset("B", 20.0)

		ap := NewAssetPrice()
		ap.AddAssetPrice("A", 2.0)
		ap.AddAssetPrice("B", 1.0)
		ap.AddAssetPrice("C", 100.0)

		got := USDValueOfEach(ua, ap)
		want := map[string]float64{"A": 20.0, "B": 20.0}

		if !reflect.DeepEqual(got, want) {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("UserAssetsTotalValue", func(t *testing.T) {
		ua := NewUserAssets()
		ua.AddAsset("A", 10.0)
		ua.AddAsset("B", 20.0)

		ap := NewAssetPrice()
		ap.AddAssetPrice("A", 2.0)
		ap.AddAssetPrice("B", 1.0)
		ap.AddAssetPrice("C", 100.0)

		got := UserAssetsTotalValue(ua, ap)
		want := 40.0

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})
}

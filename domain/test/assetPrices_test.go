package test

import (
	"testing"

	. "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

func TestAssetPrices(t *testing.T) {

	t.Run("add new asset price", func(t *testing.T) {
		ap := NewAssetPrice()

		ap.AddAssetPrice("A", 10.0)

		got := ap["A"]
		want := 10.0

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("update existing asset price", func(t *testing.T) {
		ap := NewAssetPrice()

		ap.AddAssetPrice("A", 10.0)
		// price updating
		ap.AddAssetPrice("A", 5.0)

		got := ap["A"]
		want := 5.0

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})
}

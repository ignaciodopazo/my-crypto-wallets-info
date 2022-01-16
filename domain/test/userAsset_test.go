package test

import (
	"testing"

	. "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

func TestUserAsset(t *testing.T) {

	t.Run("AddAsset", func(t *testing.T) {
		ua := NewUserAssets()
		ua.AddAsset("N", 10)
		ua.AddAsset("N", 20)

		got := ua["N"]
		want := 30.0

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("AssetQuantity", func(t *testing.T) {
		ua := NewUserAssets()
		ua.AddAsset("N", 30)
		ua.AddAsset("N", 20)

		got := ua.AssetQuantity("N")
		want := 50.0

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("PopAsset", func(t *testing.T) {
		ua := NewUserAssets()
		ua.AddAsset("N", 30)

		got := ua.PopAsset("N")
		want := 30.0

		if got != want {
			t.Error("got", got, ", want", want)
		}

		got =  ua.AssetQuantity("N")
		want = 0.0

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

}
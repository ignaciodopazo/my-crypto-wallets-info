package test

import (
	"testing"

	. "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

func TestDefaultURLBuilder(t *testing.T) {

	t.Run("test generateURL with no baseURL and no query params", func(t *testing.T) {
		b := NewURLBuilder()
		got := b.GenerateURL()
		want := ""

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("test generateURL with no baseURL and with query params", func(t *testing.T) {
		b := NewURLBuilder()
		b.AddQueryParam("key",  "value")
		got := b.GenerateURL()
		want := ""

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("test generateURL with baseURL and no query params", func(t *testing.T) {
		b := NewURLBuilder()
		b.SetBaseURL("baseurl")
		got := b.GenerateURL()
		want := "baseurl"

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("test generateURL with baseURL and with query params", func(t *testing.T) {
		b := NewURLBuilder()
		b.SetBaseURL("baseurl")
		b.AddQueryParam("key1", "value1")
		b.AddQueryParam("key2", "value2")
		got := b.GenerateURL()
		// non deterministic order of key-value pairs
		want1 := "baseurl?key1=value1&key2=value2"
		want2 := "baseurl?key2=value2&key1=value1"


		if got != want1 {
			if got != want2 {
				t.Error("got", got, ", want", want1, "or", want2)
			}
		}
	})
}
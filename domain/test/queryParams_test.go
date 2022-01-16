package test

import (
	"testing"

	. "github.com/ignaciodopazo/my-crypto-wallets-info/domain"
)

func TestAppendTo(t *testing.T) {
	t.Run("PopQueryParams an empty QueryParams", func(t *testing.T) {
		qp := NewQueryParams()
		got := qp.PopQueryParam("AA")
		want := ""

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("PopQueryParams a non-existent key in QueryParams", func(t *testing.T) {
		qp := NewQueryParams()
		qp.AddQueryParam("key", "value")
		got := qp.PopQueryParam("non exists")
		want := ""

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})

	t.Run("PopQueryParams a existen key in QueryParams", func(t *testing.T) {
		qp := NewQueryParams()
		qp.AddQueryParam("key1", "value1")
		qp.AddQueryParam("key2", "value2")
		qp.AddQueryParam("key3", "value3")
		got := qp.PopQueryParam("key2")
		want := "value2"

		if got != want {
			t.Error("got", got, ", want", want)
		}
	})
	// NewQueryParams and AddQueryParam are straight Go functions, it doesn't worth to test them
}

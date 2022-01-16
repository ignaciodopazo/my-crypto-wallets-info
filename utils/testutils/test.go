package testutils

import (
	"testing"
)

func ExpectedResult[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Error("got", got, ", want", want)
	}
}

func ExpectedResultSlices[T comparable](t testing.TB, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Error(got, "and", want, "have different length")
	}
	for i := range got {
		ExpectedResult(t, got[i], want[i])
	}
}

func CheckPredicateOverSlice[T comparable](t testing.TB, p func(T) bool, got []T) {
	t.Helper()
	for _, v := range got {
		if !p(v) {
			t.Error(got, "predicate is false")
		}
	}
}

package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum all multiples of 3 and 5 under 10", func(t *testing.T) {
		number := 10

		got := SumMultiplesOf3And5(number)
		want := 23

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, number)
		}
	})

	t.Run("Sum all multiples of 3 and 5 under 1000", func(t *testing.T) {
		number := 1000

		got := SumMultiplesOf3And5(number)
		want := 233168

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, number)
		}
	})
}

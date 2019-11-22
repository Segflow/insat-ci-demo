package main

import "testing"

func TestFact(t *testing.T) {
	var tests = map[uint]uint{
		0: 1,
		1: 1,
		5: 120,
	}

	for n, r := range tests {
		a := Fact(n)
		if a != r {
			t.Fatalf("Fact(%d) should be %d got %d", n, r, a)
		}
	}
}

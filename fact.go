package main

func Fact(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

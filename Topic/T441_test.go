package main

import "testing"

func arrangeCoins(n int) int {
	num := 0

	for n > num +1 {
		n -= num
		num++
	}

	return num
}

func TestT441(t *testing.T) {

}

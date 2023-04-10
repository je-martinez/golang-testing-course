package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("With %d got true, but is expected false", 0)
	}

	if msg != "0 is not prime by definition" {
		t.Errorf("wrong message returned: %s", msg)
	}
}

package main

import "testing"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
	}

	for _, test := range primeTests {
		result, msg := isPrime(test.testNum)
		if test.expected && !result {
			t.Errorf("%s expected true but got false", test.name)
		}

		if !test.expected && result {
			t.Errorf("%s expected false but got true", test.name)
		}

		if test.msg != msg {
			t.Errorf("%s, expected: %s but got %s", test.name, test.msg, msg)
		}
	}

}

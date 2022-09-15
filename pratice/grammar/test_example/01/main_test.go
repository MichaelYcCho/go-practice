package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if mySum(2, 3) != 5 {
		t.Error("Expected", 5, "Got", x)
	}

}

// go test -v

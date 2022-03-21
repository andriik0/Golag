package main

import "testing"

func TestMultiple(t *testing.T) {
	var x, y, result = 2, 3, 6
	funcResult := Multiple(x, y)

	if result != funcResult {
		t.Errorf("args %d, %d, result %d != %d expected result", x, y, funcResult, result)
	}
}

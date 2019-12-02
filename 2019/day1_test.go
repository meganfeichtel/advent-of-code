package main

import (
	"testing"
)

func TestCalcFuel(t *testing.T) {
	//unit tests from the website tested below
	if calcFuel(14) != 2 {
		t.Errorf("Expected fuel amount of %v, but got %v", 2, calcFuel(14))
	}
	if calcFuel(1969) != 966 {
		t.Errorf("Expected fuel amount of %v, but got %v", 966, calcFuel(1969))
	}
	if calcFuel(100756) != 50346 {
		t.Errorf("Expected fuel amount of %v, but got %v", 50346, calcFuel(100756))
	}

}

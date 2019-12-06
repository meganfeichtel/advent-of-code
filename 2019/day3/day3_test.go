package main

import (
	"testing"
)

func TestGenerateGap(t *testing.T) {
	//unit tests from the website tested below
	t1, _ := generateGAP("1,0,0,0,99")
	e1 := "2,0,0,0,99"
	if t1 != e1 {
		t.Errorf("Expected output of %s, but got %s", e1, t1)
	}
}

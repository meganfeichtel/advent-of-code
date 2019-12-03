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

	t2, _ := generateGAP("2,3,0,3,99")
	e2 := "2,3,0,6,99"
	if t2 != e2 {
		t.Errorf("Expected output of %s, but got %s", e2, t2)
	}

	t3, _ := generateGAP("2,4,4,5,99,0")
	e3 := "2,4,4,5,99,9801"
	if t3 != e3 {
		t.Errorf("Expected output of %s, but got %s", e3, t3)
	}

	t4, _ := generateGAP("1,1,1,4,99,5,6,0,99")
	e4 := "30,1,1,4,2,5,6,0,99"
	if t4 != e4 {
		t.Errorf("Expected output of %s, but got %s", e4, t4)
	}

	t5, _ := generateGAP("1,9,10,3,2,3,11,0,99,30,40,50")
	e5 := "3500,9,10,70,2,3,11,0,99,30,40,50"
	if t5 != e5 {
		t.Errorf("Expected output of %s, but got %s", e5, t5)
	}

}

func TestFindGAPInstPair(t *testing.T) {
	//unit tests from the website tested below

	//possible to find the number
	n1, v1, _ := findGAPInstPair("1,0,0,0,99", 2)
	en1 := 0
	ev1 := 0
	if n1 != en1 || v1 != ev1 {
		t.Errorf("Expected output of %v %v, but got %v %v", n1, v1, en1, ev1)
	}

	//not possible to find
	n2, v2, err := findGAPInstPair("1,0,0,0,99", 4)
	if err == nil {
		t.Errorf("Expected to get a failure but got %v %v instead", n2, v2)
	}

	//possible to find the number
	n3, v3, _ := findGAPInstPair("1,1,1,3,2,3,0,0,99", 3)
	en3 := 3
	ev3 := 0
	if n3 != en3 || v3 != ev3 {
		t.Errorf("Expected output of %v %v, but got %v %v", n3, v3, en3, ev3)
	}

	// //not possible to find
	n4, v4, err := findGAPInstPair("1,1,1,3,2,3,0,0,99", 4)
	if err == nil {
		t.Errorf("Expected to get a failure but got %v %v instead", n4, v4)
	}

	n5, v5, err := findGAPInstPair("1,1,1,3,2,3,0,0,99,1,1,1,3,1,1,1,3", 4)
	if err == nil {
		t.Errorf("Expected to get a failure but got %v %v instead", n5, v5)
	}
}

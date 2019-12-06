package main

import (
	"fmt"
	"io/ioutil"
)

var filePath = "/Users/mfeichtel/go/src/github.com/meganfeichtel/advent-of-code/2019/day3/day3.txt"

func generateGAP() {

}

func main() {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Here %v\n", f)
}

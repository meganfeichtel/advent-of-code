package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var filePath = "/Users/mfeichtel/go/src/github.com/meganfeichtel/advent-of-code/2019/day3/day3.txt"

func main() {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	wires := strings.Split(string(f), "\n")
	w1, w2 := wires[0], wires[1]

	res, _ := findShortestWireCrossingDistance(w1, w2)

	fmt.Println("Distance to shortest wire is: ", res)
}

func findShortestWireCrossingDistance(w1 string, w2 string) (int, error) {
	return -1, nil
}

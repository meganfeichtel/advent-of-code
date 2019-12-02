package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var filePath = "/Users/mfeichtel/go/src/github.com/meganfeichtel/advent-of-code/2019/inputs/day1.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcFuel(m int) int {
	f := (m / 3) - 2
	if f < 0 {
		return 0
	}
	return f + calcFuel(f)
}

func calcTotals() int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalFuel int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // internally, it advances token based on sperator
		var mass = scanner.Text() // token in unicode-char

		//convert mass to int, check errors
		massI, err := strconv.Atoi(mass)
		if err != nil {
			log.Fatal(err)
		}

		f := calcFuel(massI)
		totalFuel += f
	}

	return totalFuel
}

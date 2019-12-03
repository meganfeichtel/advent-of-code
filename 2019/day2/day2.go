package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var filePath = "/Users/mfeichtel/go/src/github.com/meganfeichtel/advent-of-code/2019/day2/day2.txt"

func convSliceStrToInt(a []string) []int {
	var b []int
	//convert array from string values to int values
	for i := 0; i < len(a); i++ {
		v1, err := strconv.Atoi(a[i])
		if err != nil {
			log.Fatal(err)
		}
		b = append(b, v1)
	}
	return b
}

//generate the gravity assist program
func generateGAP(s string) (string, error) {
	a := strings.Split(s, ",")
	b := convSliceStrToInt(a)

	//loop through the array to calculate new code instructions
	for i := 0; i < len(b)-3; i = i + 4 {
		opcode := b[i]
		if opcode == 99 {
			return strings.Trim(strings.Join(strings.Split(fmt.Sprint(b), " "), ","), "[]"), nil
		} else if opcode == 1 {
			//add the values at the indices of the next 2 values
			b[b[i+3]] = b[b[i+1]] + b[b[i+2]]
		} else if opcode == 2 {
			b[b[i+3]] = b[b[i+1]] * b[b[i+2]]
		} else {
			return "", fmt.Errorf("not an opcode! expected: 99, 1, or 2; got %v", opcode)
		}
	}
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(b), " "), ","), "[]"), nil
}

func findGAPInstPair(s string, num int) (int, int, error) {
	a := strings.Split(s, ",")
	b := convSliceStrToInt(a)

	for i := 0; i < len(b)-3; {
		opcode := b[i]
		if opcode == 1 {
			v1 := b[i+1]
			v2 := b[i+2]
			if b[v1]+b[v2] == num && b[i+3] == 0 {
				return v1, v2, nil
			}
			i = i + 4
		} else if opcode == 2 {
			v1 := b[i+1]
			v2 := b[i+2]
			if b[v1]*b[v2] == num && b[i+3] == 0 {
				return v1, v2, nil
			}
			i = i + 4
		} else if opcode == 99 {
			i = i + 1
		} else {
			return -1, -1, fmt.Errorf("never found the number %v", num)
		}
	}
	return -1, -1, fmt.Errorf("never found the number %v", num)
}

func main() {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	//pt1
	r, err := generateGAP(strings.TrimSpace(string(f)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1:\n %s\n\n", r)

	//pt2
	n, v, err := findGAPInstPair(strings.TrimSpace(string(f)), 19690720)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2:\n Noun: %v Verb: %v\n", n, v)
}

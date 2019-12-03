package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var filePath = "/Users/mfeichtel/go/src/github.com/meganfeichtel/advent-of-code/2019/day2/day2.txt"

//generate the gravity assist program
func generateGAP(s string) (string, error) {
	a := strings.Split(s, ",")
	var b []int

	for i := 0; i < len(a); i++ {
		v1, err := strconv.Atoi(a[i])
		if err != nil {
			log.Fatal(err)
		}
		b = append(b, v1)
	}

	for i := 0; i < len(b); i = i + 4 {
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

func main() {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r, err := generateGAP(strings.TrimSpace(string(f)))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

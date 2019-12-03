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

func generateGAPHelper(arr []int) ([]int, error) {
	//loop through the array to calculate new code instructions
	for i := 0; i < len(arr)-3; i = i + 4 {
		opcode := arr[i]
		if opcode == 99 {
			return arr, nil
		} else if opcode == 1 {
			//add the values at the indices of the next 2 values
			arr[arr[i+3]] = arr[arr[i+1]] + arr[arr[i+2]]
		} else if opcode == 2 {
			//multiply the values at the indices of the next 2 values
			arr[arr[i+3]] = arr[arr[i+1]] * arr[arr[i+2]]
		} else {
			return []int{}, fmt.Errorf("not an opcode! expected: 99, 1, or 2; got %v", opcode)
		}
	}
	return arr, nil
}

//generate the gravity assist program
func generateGAP(s string) (string, error) {
	a := strings.Split(s, ",")
	b := convSliceStrToInt(a)
	c, err := generateGAPHelper(b)
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(c), " "), ","), "[]"), err
}

func findGAPInstPairHelper(arr []int, num int) (int, int, error) {
	for i := 0; i < 100; i++ {
		arr[1] = i

		for j := 0; j < 100; j++ {
			arr[2] = j
			cp := make([]int, len(arr))
			copy(cp, arr)
			res, _ := generateGAPHelper(cp)

			if res[0] == num {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("never found the number %v", num)
}

func findGAPInstPair(s string, num int) (int, int, error) {
	a := strings.Split(s, ",")
	b := convSliceStrToInt(a)
	return findGAPInstPairHelper(b, num)
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

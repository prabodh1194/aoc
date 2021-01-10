package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var nums []int

func main() {
	file, err := os.Open("d1.inp")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	a, b := sumN(2020)
	log.Printf("Found number: %d", a*b)
}

func sumN(fnum int) (int, int) {
	for _, num := range nums {
		a, b := sum2(fnum - num)

		if a != -1 {
			return a * b, num
		}
	}

	return -1, -1
}

func sum2(fnum int) (int, int) {
	var m map[int]bool
	m = make(map[int]bool)

	for _, num := range nums {
		if m[fnum-num] {
			return num, (fnum - num)
		} else {
			m[num] = true
		}
	}

	return -1, -1
}

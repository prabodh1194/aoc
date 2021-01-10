package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("d1.inp")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	m := make(map[int]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		if m[2020-num] {
			log.Printf("Found a number %d", num*(2020-num))
		} else {
			m[num] = true
		}
	}
}

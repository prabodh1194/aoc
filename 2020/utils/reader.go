package utils

import (
	"bufio"
	"log"
	"os"
)

func FileReader(path string) <-chan string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	stringChnl := make(chan string)

	go func() {
		for scanner.Scan() {
			stringChnl <- scanner.Text()
		}
		close(stringChnl)
		file.Close()
	}()

	return stringChnl
}

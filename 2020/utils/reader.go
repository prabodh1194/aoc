package utils

import (
	"bufio"
	"log"
	"os"
)

// FileReader reads file and returns line-by-line.
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

// FileReaderInGroups reads file till it encounters a blank line. All the text that has been read till now will be
// returned as a per-line array.
func FileReaderInGroups(path string) <-chan []string {
	var txts []string
	stringChnl := make(chan []string)

	go func() {
		for txt := range FileReader(path) {
			if len(txt) == 0 {
				stringChnl <- txts
				txts = make([]string, 0)
			} else {
				txts = append(txts, txt)
			}
		}
		stringChnl <- txts
		close(stringChnl)
	}()

	return stringChnl
}

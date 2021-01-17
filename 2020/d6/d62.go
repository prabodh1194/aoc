package d6

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func D62() {
	var m map[string]int
	sum := 0

	for txts := range utils.FileReaderInGroups("d6/d6.inp") {
		m = make(map[string]int)

		for _, txt := range txts {
			for _, t := range strings.Split(txt, "") {
				m[t]++
			}
		}
		fmt.Println(m, len(txts))

		for _, v := range m {
			if v == len(txts) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

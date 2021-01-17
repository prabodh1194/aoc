package d6

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func D61() {
	var m map[string]struct{}
	sum := 0

	for txts := range utils.FileReaderInGroups("d6/d6.inp") {
		m = make(map[string]struct{})

		for _, txt := range txts {
			for _, t := range strings.Split(txt, "") {
				m[t] = struct{}{}
			}
		}

		fmt.Println(len(m))
		sum += len(m)
	}

	fmt.Println(sum)
}

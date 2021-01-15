package d3

import (
	"aoc/utils"
	"fmt"
)

func D31() {
	for txt := range utils.FileReader("d3/d3.inp") {
		fmt.Println(txt)
	}
}

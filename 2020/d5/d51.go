package d5

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func D51() {
	max := int64(0)
	for txt := range utils.FileReader("d5/d5.inp") {
		rs := txt[0:7]
		cs := txt[7:]

		cs = strings.ReplaceAll(cs, "L", "0")
		cs = strings.ReplaceAll(cs, "R", "1")
		rs = strings.ReplaceAll(rs, "F", "0")
		rs = strings.ReplaceAll(rs, "B", "1")

		r, _ := strconv.ParseInt(rs, 2, 64)
		c, _ := strconv.ParseInt(cs, 2, 64)

		id := r*8 + c

		if id > max {
			max = id
		}

		fmt.Println(rs, cs, r, c, id)
	}

	fmt.Println(max)
}

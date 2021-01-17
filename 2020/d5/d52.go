package d5

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func D52() {
	var m map[int64]bool
	m = make(map[int64]bool)

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
		m[id] = true
	}

	i := int64(0)
	for ; i < 1024; i++ {
		if !m[i] && m[i+1] && m[i-1] {
			fmt.Println(i)
		}
	}
}

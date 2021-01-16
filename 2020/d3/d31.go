package d3

import (
	"aoc/utils"
	"fmt"
)

func D31() {
	var frst [][]byte
	for txt := range utils.FileReader("d3/d3.inp") {
		var row []byte
		for i := 0; i < len(txt); i++ {
			row = append(row, txt[i])
		}
		frst = append(frst, row)
	}

	cnt := 1

	cnt *= getSlope(frst, 1, 1)
	cnt *= getSlope(frst, 1, 3)
	cnt *= getSlope(frst, 1, 5)
	cnt *= getSlope(frst, 1, 7)
	cnt *= getSlope(frst, 2, 1)

	fmt.Println(cnt)
}

func getSlope(frst [][]byte, ds int, rs int) int {
	cnt := 0

	for i, j := 0, 0; i < len(frst); i, j = i+ds, (j+rs)%len(frst[i]) {
		if frst[i][j] == 35 {
			cnt++
		}
	}
	return cnt
}

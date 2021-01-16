package d4

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func D41() {
	var passport []string
	cnt := 0
	for txt := range utils.FileReader("d4/d4.inp") {
		if len(txt) == 0 {
			if validatePassport(passport) {
				cnt++
			}
			passport = make([]string, 0)
		} else {
			passport = append(passport, strings.Split(txt, " ")...)
		}
	}

	if validatePassport(passport) {
		cnt++
	}

	fmt.Println(cnt)
}

func validatePassport(passport []string) bool {
	cnt := 0
	cidAbsent := true
	flag := false

	for _, attrib := range passport {
		cnt++
		k := strings.Split(attrib, ":")[0]

		if k == "cid" {
			cidAbsent = false
		}
	}

	if cnt == 7 && cidAbsent {
		flag = true
	}

	if cnt == 8 {
		flag = true
	}

	return flag
}

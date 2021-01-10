package d2

import (
	"aoc/utils"
	"log"
	"regexp"
	"strconv"
)

func D21() {
	r := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>[a-z]): (?P<pwd>[a-z]+)`)
	valid := 0

	for txt := range utils.FileReader("d2/d2.inp") {
		res := utils.RegexpToMap(r, txt)
		cc := 0
		toCatch := rune(res["char"][0])

		for _, c := range res["pwd"] {
			if c == toCatch {
				cc++
			}
		}

		min, _ := strconv.Atoi(res["min"])
		max, _ := strconv.Atoi(res["max"])

		if cc >= min && cc <= max {
			valid++
		}
	}

	log.Printf("Valid passwords are: %d", valid)
}

package d2

import (
	"aoc/utils"
	"log"
	"regexp"
	"strconv"
)

func D22() {
	r := regexp.MustCompile(`(?P<fir>\d+)-(?P<sec>\d+) (?P<char>[a-z]): (?P<pwd>[a-z]+)`)
	valid := 0

	for txt := range utils.FileReader("d2/d2.inp") {
		res := utils.RegexpToMap(r, txt)
		toCatch := res["char"][0]
		pwd := res["pwd"]

		fir, _ := strconv.Atoi(res["fir"])
		sec, _ := strconv.Atoi(res["sec"])

		fir--
		sec--

		if pwd[fir] == pwd[sec] && pwd[fir] == toCatch {
		} else if pwd[fir] == toCatch || pwd[sec] == toCatch {
			valid++
		}
	}

	log.Printf("Valid passwords are: %d", valid)
}

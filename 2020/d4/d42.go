package d4

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func D42() {
	var passport []string
	cnt := 0
	for txt := range utils.FileReader("d4/d4.inp") {
		if len(txt) == 0 {
			if validatePassportWithCreds(passport) {
				cnt++
			}
			passport = make([]string, 0)
		} else {
			passport = append(passport, strings.Split(txt, " ")...)
		}
	}

	if validatePassportWithCreds(passport) {
		cnt++
	}

	fmt.Println(cnt)
}

func validatePassportWithCreds(passport []string) bool {
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
	} else if cnt == 8 {
		flag = true
	}

	for _, attrib := range passport {
		if !flag {
			return flag
		}

		k := strings.Split(attrib, ":")

		switch k[0] {
		case "byr":
			flag = validateByr(k[1])
		case "iyr":
			flag = validateIyr(k[1])
		case "eyr":
			flag = validateEyr(k[1])
		case "hgt":
			flag = validateHgt(k[1])
		case "hcl":
			flag = validateHcl(k[1])
		case "ecl":
			flag = validateEcl(k[1])
			if flag {
				fmt.Println(k[1])
			}
		case "pid":
			flag = validatePid(k[1])
		}
	}

	return flag
}

func validateYr(s string, lo, hi int) bool {
	r := regexp.MustCompile(`^\d{4}$`)
	if !r.MatchString(s) {
		return false
	}

	yr, _ := strconv.Atoi(s)

	return yr >= lo && yr <= hi
}

func validateByr(s string) bool {
	return validateYr(s, 1920, 2002)
}

func validateIyr(s string) bool {
	return validateYr(s, 2010, 2020)
}

func validateEyr(s string) bool {
	return validateYr(s, 2020, 2030)
}

func validateHgt(s string) bool {
	r := regexp.MustCompile(`(?m)(?P<hgtin>^\d{2})in$|(?P<hgtcm>^\d{3})cm$`)
	m := utils.RegexpToMap(r, s)

	if m["hgtcm"] != "" {
		hgt, _ := strconv.Atoi(m["hgtcm"])
		return hgt >= 150 && hgt <= 193
	} else if m["hgtin"] != "" {
		hgt, _ := strconv.Atoi(m["hgtin"])
		return hgt >= 59 && hgt <= 76
	}

	return false
}

func validateHcl(s string) bool {
	r := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return r.MatchString(s)
}

var validEcls = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func validateEcl(s string) bool {
	fmt.Println(validEcls[s])
	return validEcls[s] == true
}

func validatePid(s string) bool {
	r := regexp.MustCompile(`^\d{9}$`)
	return r.MatchString(s)
}

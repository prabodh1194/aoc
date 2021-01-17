package d7

import (
	"aoc/utils"
	"fmt"
	"regexp"
)

var sum = 0
var adjMap = make(map[string][]string)
var visited = make(map[string]bool)

func D71() {
	re1 := regexp.MustCompile(`(?m)(?P<color>[[:alpha:]]+ [[:alpha:]]+) bags contain`)
	re2 := regexp.MustCompile(`(?m)((?P<count>\d+)|no) (other|(?P<color>[[:alpha:]]+ [[:alpha:]]+)) bag(s|)[,.]`)

	for txt := range utils.FileReader("d7/d7.inp") {
		outBag := utils.RegexpToMap(re1, txt)["color"]

		var inBags []string
		for _, bag := range utils.MultiRegexToMap(re2, txt) {
			inBags = append(inBags, bag["color"])
		}

		for _, bag := range inBags {
			adjMap[bag] = append(adjMap[bag], outBag)
		}
	}

	fmt.Println(adjMap)
	count("shiny gold")
	fmt.Println(sum)
}

func count(bag string) {
	for _, cntBag := range adjMap[bag] {
		if !visited[cntBag] {
			sum++
			visited[cntBag] = true
		}
		count(cntBag)
	}
}

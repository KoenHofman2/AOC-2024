package day3

import (
	"aoc-2024/utils"
	"fmt"
	"regexp"
)

func PartOne() {
	solution(solutionOne)
}

func PartTwo() {
	solution(solutionTwo)
}

func solution(fn func(string) int) {
	fmt.Println(fn(utils.ReadRawTxtFile("day3/day3.txt")))
}

func solutionOne(text string) int {
	return solutionWrapper(text, `mul\([0-9]{1,3},[0-9]{1,3}\)`, false)
}

func solutionTwo(text string) int {
	return solutionWrapper(text, `mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`, true)
}

func solutionWrapper(text, regexString string, doDont bool) (sum int) {
	r := regexp.MustCompile(regexString)
	matches := r.FindAllString(text, -1)
	matchRegex := regexp.MustCompile(`[0-9]{1,3}`)
	enabled := true

	for _, match := range matches {
		if doDont {
			switch match {
			case "don't()":
				enabled = false

				continue
			case "do()":
				enabled = true

				continue
			}
		}

		if enabled {
			numbers := matchRegex.FindAllString(match, -1)
			sum += utils.StringToInt(numbers[0]) * utils.StringToInt(numbers[1])
		}
	}

	return
}

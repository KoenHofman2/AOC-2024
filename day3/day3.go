package day3

import (
	"aoc-2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func PartOne() {
	solution(solutionOne)
}

func PartTwo() {
	solution(solutionTwo)
}

func solution(fn func(string) int) {
	text := utils.ReadRawTxtFile("day3/day3.txt")

	sum := fn(text)

	fmt.Println(sum)
}

func solutionOne(text string) (sum int) {
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := r.FindAllString(text, -1)
	matchRegex := regexp.MustCompile(`[0-9]{1,3}`)

	for _, match := range matches {
		numbers := matchRegex.FindAllString(match, -1)

		firstNumber, _ := strconv.Atoi(numbers[0])
		secondNumber, _ := strconv.Atoi(numbers[1])

		sum += firstNumber * secondNumber
	}

	return
}

func solutionTwo(text string) (sum int) {
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matches := r.FindAllString(text, -1)
	matchRegex := regexp.MustCompile(`[0-9]{1,3}`)
	enabled := true

	for _, match := range matches {
		if match == "don't()" {
			enabled = false

			continue
		} else if match == "do()" {
			enabled = true

			continue
		}

		if !enabled {
			continue
		}

		numbers := matchRegex.FindAllString(match, -1)

		firstNumber, _ := strconv.Atoi(numbers[0])
		secondNumber, _ := strconv.Atoi(numbers[1])

		sum += firstNumber * secondNumber
	}

	return
}

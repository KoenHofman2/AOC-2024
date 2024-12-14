package day2

import (
	"aoc-2024/utils"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func PartOne() {
	solution(solutionOne)
}

func PartTwo() {
	solution(solutionTwo)
}

func solution(fn func([]int) bool) {
	scanner, file := utils.ReadTxtFile("day2/day2.txt")
	defer file.Close()

	var safeCount int

	for _, values := range getValues(scanner) {
		if fn(values) {
			safeCount++
		}
	}

	fmt.Println("safeCount: ", safeCount)
}

func getValues(scanner *bufio.Scanner) [][]int {
	var results [][]int

	for scanner.Scan() {
		var intValues []int

		for _, strValue := range strings.Split(scanner.Text(), " ") {
			intValue, _ := strconv.Atoi(strValue)
			intValues = append(intValues, intValue)
		}

		results = append(results, intValues)

	}

	return results
}

func solutionOne(values []int) bool {
	firstDifference := values[1] - values[0]

	for i, value := range values {
		if i == 0 {
			continue
		}

		difference := value - values[i-1]

		if math.Abs(float64(difference)) > 3 || firstDifference*difference <= 0 {
			return false
		}
	}

	return true
}

func solutionTwo(values []int) bool {
	if !solutionOne(values) {
		for i := range values {
			filteredResult := solutionOne(utils.RemoveElementAt(utils.CopySlice(values), i))
			if filteredResult {
				return true
			}
		}

		return false
	}

	return true
}

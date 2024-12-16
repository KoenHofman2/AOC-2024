package day4

import (
	"aoc-2024/utils"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func PartOne() {
	solution(solutionOne)
}

func PartTwo() {
	solution(solutionTwo)
}

func solution(fn func([]string) int) {
	scanner, file := utils.ReadTxtFile("day4/day4.txt")
	defer file.Close()

	lines := getValues(scanner)

	fmt.Println(fn(lines))
}

func getValues(scanner *bufio.Scanner) []string {
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func solutionOne(lines []string) int {
	r := regexp.MustCompile(`XMAS|SAMX`)

	var horizontalCount, verticalCount, leftDiagonalCount, rightDiagonalCount int
	var allCharacters [][]string // horizontal - vertical

	linesCount := len(lines)
	columnCount := len(lines[0])

	for _, line := range lines {
		horizontalCount += len(r.FindAllString(line, -1))

		characters := strings.Split(line, "")
		allCharacters = append(allCharacters, characters)
	}

	for rowIndex := 0; rowIndex < columnCount; rowIndex++ {
		var columnString, rightDiagonalString, leftDiagonalString string

		for columnIndex := 0; columnIndex < linesCount; columnIndex++ {
			columnString += allCharacters[columnIndex][rowIndex]

			if rowIndex-columnIndex >= 0 {
				rightDiagonalString += allCharacters[columnIndex][rowIndex-columnIndex]
			}
		}

		// todo: ADJUST THIS
		//for columnIndex := 140; columnIndex >= 0; columnIndex-- {
		//	if columnIndex-1 >= 0 && rowIndex-1 >= 0 {
		//		leftDiagonalString += allCharacters[columnIndex+1][rowIndex+1]
		//	}
		//}

		// Generate the left diagonal string (top-right to bottom-left)
		for columnIndex := 0; columnIndex < linesCount; columnIndex++ {
			// Calculate the corresponding column index for the left diagonal
			// Start from the top-right of the matrix
			leftDiagonalIndex := linesCount - 1 - columnIndex
			// Ensure the column index stays within bounds
			if leftDiagonalIndex >= 0 && rowIndex+columnIndex < linesCount {
				leftDiagonalString += allCharacters[rowIndex+columnIndex][leftDiagonalIndex]
			}
		}

		verticalCount += len(r.FindAllString(columnString, -1))
		rightDiagonalCount += len(r.FindAllString(rightDiagonalString, -1))
		leftDiagonalCount += len(r.FindAllString(leftDiagonalString, -1))
	}

	fmt.Println("horizontalCount:", horizontalCount)
	fmt.Println("verticalCount:", verticalCount)
	fmt.Println("leftDiagonalCount:", leftDiagonalCount)
	fmt.Println("rightDiagonalCount:", rightDiagonalCount)

	return horizontalCount + verticalCount + leftDiagonalCount + rightDiagonalCount
}

func solutionTwo(lines []string) int {
	return 0
}

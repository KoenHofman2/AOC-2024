package main

import (
	"aoc-2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	left, right, rightMap := leftAndRightFromFile()

	slices.Sort(left)
	slices.Sort(right)

	var distance int

	for i, _ := range left {
		distance += utils.AbsoluteDifference(left[i], right[i])
	}

	var similarityScore int

	for _, leftValue := range left {
		similarityScore += leftValue * rightMap[leftValue]
	}

	fmt.Println("distance: ", distance)
	fmt.Println("similarityScore: ", similarityScore)
	fmt.Println("elapsed  time: ", time.Since(start).Nanoseconds())

}

func leftAndRightFromFile() ([]int, []int, map[int]int) {
	scanner, file := utils.ReadTxtFile("day1/day1.txt")
	defer file.Close()

	var (
		left, right []int
		rightMap    = make(map[int]int)
	)

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, "   ")

		leftNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		rightNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftNumber)
		right = append(right, rightNumber)

		rightValue, ok := rightMap[rightNumber]
		if ok {
			rightMap[rightNumber] = rightValue + 1
		} else {
			rightMap[rightNumber] = 1
		}
	}

	return left, right, rightMap
}

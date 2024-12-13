package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	left, rightMap := leftAndRightFromFileMap()

	var similarityScore int

	for _, leftValue := range left {
		similarityScore += leftValue * rightMap[leftValue]
	}

	fmt.Println(similarityScore)
	fmt.Println(time.Since(start))

}

func leftAndRightFromFileMap() ([]int, map[int]int) {
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		left     []int
		rightMap = make(map[int]int)
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

		rightValue, ok := rightMap[rightNumber]
		if ok {
			rightMap[rightNumber] = rightValue + 1
		} else {
			rightMap[rightNumber] = 1
		}
	}

	return left, rightMap
}

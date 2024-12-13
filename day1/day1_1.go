package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	left, right := leftAndRightFromFile()

	slices.Sort(left)
	slices.Sort(right)

	var distance int

	for i, _ := range left {
		distance += absDiff(left[i], right[i])
	}

	fmt.Println(distance)
	fmt.Println(time.Since(start).Nanoseconds())

}

func leftAndRightFromFile() ([]int, []int) {
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left, right []int

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
	}

	return left, right
}

package utils

import (
	"bufio"
	"os"
)

func ReadTxtFile(filePath string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file), file
}

func AbsoluteDifference(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func RemoveElementAt(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func CopySlice(values []int) []int {
	valuesCopy := make([]int, len(values))
	copy(valuesCopy, values[:])

	return valuesCopy
}

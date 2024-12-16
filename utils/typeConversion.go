package utils

import "strconv"

func StringToInt(strValue string) int {
	intValue, _ := strconv.Atoi(strValue)

	return intValue
}

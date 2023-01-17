package common

import (
	"fmt"
	"strconv"
)

func FormattedOutput(text interface{}) {
	fmt.Println(text)
}

func ConvertToInteger(num string) int {
	numInt, _ := strconv.Atoi(num)
	return numInt
}

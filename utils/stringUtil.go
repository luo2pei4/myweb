package utils

import (
	"fmt"
	"strconv"
)

// StringToInt 将string转换为int, 转换异常的情况下返回0
func StringToInt(str string) int {

	value, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return value
}

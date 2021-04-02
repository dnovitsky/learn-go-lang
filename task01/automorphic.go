package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &num)

	squareNum := square(num)
	is := isAutomorphic(num, squareNum)

	if is {
		fmt.Print("Is it automorphic? Yes")
	} else {
		fmt.Print("Is it automorphic? No")
	}
}

func square(num int) int {
	return num * num
}

func isAutomorphic(num int, sNum int) bool {
	numStr := strconv.Itoa(num)
	sNumStr := strconv.Itoa(sNum)

	return strings.HasSuffix(sNumStr, numStr)
}

package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var (
		option   string
		inputStr string
	)

	fmt.Print("Select option: compression or decompression? Enter character 'C'/'D': ")
	fmt.Scan(&option)

	switch option {
	case "C":
		inputStr = inputString()
		inputStr = compress(inputStr)
		fmt.Printf("The string has been compressed: %s", inputStr)
		break
	case "D":
		inputStr = inputString()
		inputStr = decompress(inputStr)
		fmt.Printf("The string has been decompressed: %s", inputStr)
		break
	default:
		fmt.Print("This option is not available.")
		break
	}
}

func inputString() string {
	var str string
	fmt.Print("Enter string: ")
	fmt.Scan(&str)
	return str
}

func compress(str string) string {
	var (
		prevS         string = ""
		currentStr    string = ""
		compressedStr string = ""
	)

	for _, v := range str {
		var s = string(v)
		if prevS == s {
			currentStr = currentStr + s
		} else {
			compressedStr = compressionStep(currentStr, compressedStr, prevS)
			prevS = s
			currentStr = s
		}
	}
	compressedStr = compressionStep(currentStr, compressedStr, prevS)
	return compressedStr
}

func compressionStep(curStr string, comprStr string, prevS string) string {
	var (
		resultStr string = ""
		len       int    = 0
	)

	len = utf8.RuneCountInString(curStr)

	if len <= 4 {
		resultStr = comprStr + curStr
	} else {
		resultStr = comprStr + "#" + strconv.Itoa(len) + "#" + prevS
	}

	return resultStr
}

func decompress(str string) string {
	for i := range str {
		var s = string(str[i])
		if s == "a" {
			fmt.Println(s)
		}
	}
	return str
}

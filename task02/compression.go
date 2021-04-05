package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
	var (
		state           int    = 0
		number          int    = 0
		currentStr      string = ""
		decompressedStr string = ""
	)

	for _, v := range str {
		var s = string(v)

		switch state {
		case 1:
			if s == "#" {
				state = 1
				decompressedStr = decompressedStr + currentStr
				currentStr = s
			} else {
				validateAndProcessNumberStep(&state, &number, &currentStr, &decompressedStr, v, s)
			}
			break
		case 2:
			if s == "#" {
				state = 3
				currentStr = currentStr + s
			} else {
				validateAndProcessNumberStep(&state, &number, &currentStr, &decompressedStr, v, s)
			}
			break
		case 3:
			decompressedStr = decompressedStr + strings.Repeat(s, number)
			setStateZeroStep(&state, &number, &currentStr)
			break
		default:
			firstSpecialCharOrDefaultStep(&state, &currentStr, &decompressedStr, s)
			break
		}
	}
	decompressedStr = decompressedStr + currentStr
	return decompressedStr
}

func firstSpecialCharOrDefaultStep(state *int, curStr *string, decomprStr *string, s string) {
	if s == "#" {
		*state = 1
		*curStr = s
	} else {
		*decomprStr = *decomprStr + s
	}
}

func validateAndProcessNumberStep(state *int, number *int, curStr *string, decomprStr *string, v rune, s string) {
	if unicode.IsDigit(v) {
		n, err := strconv.Atoi(s)
		if err == nil {
			*state = 2
			*number = *number * 10
			*number = *number + n
			*curStr = *curStr + s
		} else {
			completeCurrentStep(curStr, decomprStr, s)
			setStateZeroStep(state, number, curStr)
		}
	} else {
		completeCurrentStep(curStr, decomprStr, s)
		setStateZeroStep(state, number, curStr)
	}
}

func completeCurrentStep(curStr *string, decomprStr *string, s string) {
	*curStr = *curStr + s
	*decomprStr = *decomprStr + *curStr
}

func setStateZeroStep(state *int, number *int, curStr *string) {
	*state = 0
	*number = 0
	*curStr = ""
}

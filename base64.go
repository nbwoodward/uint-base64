package base64

import (
	"errors"
	"fmt"
)

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
var intMap = make(map[string]int)
var charMap = make(map[int]string)

func init() {
	setupMaps()
}

func IntToString(src int) (hash string, err error) {
	if src < 0 {
		return "", errors.New("Integer must be >= 0")
	}

	for {
		key := src & 63
		hash = string(charMap[key]) + hash

		src = src >> 6
		if src == 0 {
			break
		}

	}

	return
}

func StringToInt(str string) (res int, err error) {
	for {
		if len(str) == 0 {
			break
		}

		char := str[0:1]
		str = str[1:]
		num, exists := intMap[char]
		if !exists {
			return 0, errors.New(fmt.Sprintf("Illegal character in string: %v", char))
		}
		res = res<<6 | num
	}

	return
}

func setupMaps() {
	for i := 0; i < len(chars); i++ {
		char := chars[i : i+1]
		intMap[char] = i
		charMap[i] = char
	}
}

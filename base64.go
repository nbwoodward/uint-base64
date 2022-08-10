package base64

import (
	"errors"
	"fmt"
)

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
var intMap = make(map[string]uint)
var charMap = make(map[uint]string)

func init() {
	setupMaps()
}

func IntToString(src uint) (hash string, err error) {
	for {
		key := src & 63
		hash = charMap[key] + hash

		src = src >> 6
		if src == 0 {
			break
		}

	}

	return
}

func StringToInt(str string) (res uint, err error) {
	if str == "" {
		return 0, errors.New("Cannot convert empty string")
	}

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
		intMap[char] = uint(i)
		charMap[uint(i)] = char
	}
}

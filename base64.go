package main

import (
	"fmt"
	"math"
)

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
var intMap = make(map[string]int)
var charMap = make(map[int]string)

func init() {
	setupMaps()
}

func main() {
	fmt.Println(math.MaxInt64)
	fmt.Println(IntToString(math.MaxInt64))
	fmt.Println(IntToString(math.MinInt64))
}

func IntToString(src int) (hash string) {
	for {
		key := src & 63
		hash = string(charMap[key]) + hash

		src = src >> 6
		if src == 0 {
			break
		}

	}

	return hash
}

func StringToInt(str string) (res int) {
	for {
		if len(str) == 0 {
			break
		}

		char := str[0:1]
		str = str[1:]
		num := intMap[char]
		res = res<<6 | num
	}

	return res
}

func setupMaps() {
	for i := 0; i < len(chars); i++ {
		char := chars[i : i+1]
		intMap[char] = i
		charMap[i] = char
	}
}

package main

import (
	"math"
	"testing"
)

var tests = []struct {
	intr int
	str  string
}{
	{0, "0"},
	{1, "1"},
	{2, "2"},
	{3, "3"},
	{4, "4"},
	{5, "5"},
	{6, "6"},
	{7, "7"},
	{8, "8"},
	{10, "a"},
	{11, "b"},
	{12, "c"},
	{13, "d"},
	{14, "e"},
	{15, "f"},
	{16, "g"},
	{17, "h"},
	{18, "i"},
	{19, "j"},
	{20, "k"},
	{21, "l"},
	{22, "m"},
	{23, "n"},
	{24, "o"},
	{25, "p"},
	{26, "q"},
	{27, "r"},
	{28, "s"},
	{29, "t"},
	{30, "u"},
	{31, "v"},
	{32, "w"},
	{33, "x"},
	{34, "y"},
	{35, "z"},
	{36, "A"},
	{37, "B"},
	{38, "C"},
	{39, "D"},
	{40, "E"},
	{41, "F"},
	{42, "G"},
	{43, "H"},
	{44, "I"},
	{45, "J"},
	{46, "K"},
	{47, "L"},
	{48, "M"},
	{49, "N"},
	{50, "O"},
	{51, "P"},
	{52, "Q"},
	{53, "R"},
	{54, "S"},
	{55, "T"},
	{56, "U"},
	{57, "V"},
	{58, "W"},
	{59, "X"},
	{60, "Y"},
	{61, "Z"},
	{62, "-"},
	{64, "10"},
	{65, "11"},
	{66, "12"},
	{127, "1_"},
	{1023, "f_"},
	{1024, "g0"},
	{1025, "g1"},
	{4095, "__"},
	{4096, "100"},
	{4097, "101"},
	{314159, "1cIL"},
	{3141592, "b-_o"},
	{31415926, "1TRVS"},
	{314159265, "iKr2x"},
	{3141592653, "2Xgepd"},
	{31415926535, "tgyfY7"},
	{314159265359, "4ABmvpf"},
	{int(math.MaxInt64), "7__________"},
}

func TestIntToString(t *testing.T) {
	for _, tt := range tests {
		result := IntToString(tt.intr)
		if result != tt.str {
			t.Errorf("Got %v, wanted %v", result, tt.str)
		}
	}
}

func TestStringToInt(t *testing.T) {
	for _, tt := range tests {
		result := StringToInt(tt.str)
		if result != tt.intr {
			t.Errorf("Got %v, wanted %v", result, tt.intr)
		}
	}

}

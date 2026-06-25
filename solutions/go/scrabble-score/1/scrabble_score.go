package scrabblescore

import (
	"slices"
	"strings"
)

func value(char string) int {
	char = strings.ToUpper(char)
	switch {
	case slices.Contains([]string{"A", "E", "I", "O", "U", "L", "N", "R", "S" , "T"}, char):
		return 1
	case slices.Contains([]string{"D", "G"}, char):
		return 2
	case slices.Contains([]string{"B", "C", "M", "P"}, char):
		return 3
	case slices.Contains([]string{"F", "H", "V", "W", "Y"}, char):
		return 4
	case char == "K":
		return 5
	case slices.Contains([]string{"J", "X"}, char):
		return 8
	case slices.Contains([]string{"Q", "Z"}, char):
		return 10
	}
	return 0
}

func Score(word string) int {
	word = strings.ToUpper(word)
	var sum int

	for i:=0 ; i<len(word);i++ {
		if word[i] < 'A' || word[i] >'Z' {
			continue
		}

		sum += value(string(word[i]))
	}
	return sum

}

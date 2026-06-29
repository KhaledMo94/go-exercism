package etl

import (
	"strings"
	// "unicode"
)

func Transform(in map[int][]string) map[string]int {
	r := make(map[string]int)

	for i , slice := range in {
		for _ , char := range slice {
			r[strings.ToLower(char)] = i
		}
	}

	return  r
}

package bottlesong

import (
	"fmt"
	"strings"
)

var words = []string{
	"no",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
}

func bottle(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}


func verse(n int) []string {
	return []string{
		fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(words[n]), bottle(n)),
		fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(words[n]), bottle(n)),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s green %s hanging on the wall.", words[n-1], bottle(n-1)),
	}
}

func Recite(startBottles, takeDown int) []string {
	var result []string

	for i := startBottles; i > startBottles-takeDown; i-- {
		result = append(result, verse(i)...)

		if i != startBottles-takeDown+1 {
			result = append(result, "")
		}
	}

	return result
}
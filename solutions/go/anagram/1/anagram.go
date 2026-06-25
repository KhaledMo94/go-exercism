package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {

	var result []string
	subject = strings.ToUpper(subject)

	
	for i, _ := range candidates {
		counts := map[rune]int{}
	
		for _ , r := range subject {
			counts[r]++
		}
		item := strings.ToUpper(candidates[i])
		if subject == item|| len(subject) != len(item){
			continue
		}

		e := true 
		for _,char := range item{
			counts[char]--
		}

		for _, count := range counts {
			if count != 0 {
				e = false
			}
		}

		if e{
			result = append(result, candidates[i])
		}
	} 

	return  result

}

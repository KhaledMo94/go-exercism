package acronym

import "strings"

func Abbreviate(s string) string {
	replaced := strings.ReplaceAll(s , "-", " ")
	replaced = strings.ReplaceAll(replaced , "_", " ")

	parts := strings.Fields(replaced)

	result := []rune{}

	for _ , part := range parts {
		result = append(result, rune(part[0]))
	}
	return  strings.ToUpper(string(result))

}

package atbashcipher

import (
	"strings"
	"unicode"
)

func mapped(char byte) byte {
	return  'z' - (char - 'a')
} 

func appendAndIncreaseCounter(result[]byte , char byte , counter *uint8) []byte{
	result = append(result, char)
	(*counter)++
	if(*counter == 5){
		result = append(result, ' ')
		*counter = 0 
	}
	return result
}

func Atbash(s string) string {
	result := []byte{}	
	var counter uint8

	for i:=0 ; i<len(s) ; i++ {
		if unicode.IsPunct(rune(s[i])) || 
			unicode.IsSpace(rune(s[i])){
			continue
		}

		if unicode.IsDigit(rune(s[i])) {
			result = appendAndIncreaseCounter(result , s[i] , &counter )
		}
		if unicode.IsLetter(rune(s[i])) {
			r := byte(unicode.ToLower(rune(s[i])))
			result = appendAndIncreaseCounter(result, mapped(r), &counter)
		}
	}

	if len(s) == len(string(result)) {
		return string(result)
	}

	return  strings.TrimSpace(string(result))

}

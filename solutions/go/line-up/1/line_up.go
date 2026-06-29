package lineup

import (
	"strconv"
	"strings"
)

func Format(name string, number int) string {
	sentance := ":Name, you are the :NumberWithSuffix customer we serve today. Thank you!"

	isTh := false
	s := strconv.Itoa(number)
	l := len(s)

	if number > 10 {
		if (s[l-2] == '1'){
			isTh = true
		}
	}

	sentance = strings.Replace(sentance,":Name",name,1)

	if isTh {
		return  strings.Replace(sentance ,":NumberWithSuffix",string(s + "th"),1 )
	}

	suffix := ""
	if(s[l-1] == '1'){
		suffix = "st"
	}else if s[l-1] == '2' {
		suffix = "nd"
	}else if s[l-1] == '3'{
		suffix = "rd"
	}else{
		suffix = "th"
	}	
	return  strings.Replace(sentance ,":NumberWithSuffix",string(s + suffix),1 )


}

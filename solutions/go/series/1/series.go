package series

// import "strings"

func All(n int, s string) []string {
    if n > len(s) || n <1 {
        return []string{}
    }

    var res []string

    for i := range len(s)-n+1 {
        res = append(res , s[i : i+n])
    }

    return res
}

func UnsafeFirst(n int, s string) string {
	if len(s) < n {
		return ""
	}

	return s[0:n]
}

package main

import (
	"strconv"
	"strings"
)

func upperWordPreserveFormat(s string) string {
	prefix := ""
	suffix := ""
	for len(s) > 0 {
		if strings.ContainsRune("([", rune(s[0])) {
			prefix += string(s[0])
			s = s[1:]
			continue
		}
		break
	}
	for len(s) > 0 {
		last := rune(s[len(s)-1])
		if strings.ContainsRune("),]", last) || last == ',' || last == ';' || last == ':' {
			suffix = string(last) + suffix
			s = s[:len(s)-1]
			continue
		}
		break
	}
	return prefix + strings.ToUpper(s) + suffix
}

func up(words []string) []string {
	var result []string

	for _, w := range words {
		normalized := strings.Trim(w, "[](),")

		switch {
		case normalized == "up":
			if len(result) > 0 {
				result[len(result)-1] = upperWordPreserveFormat(result[len(result)-1])
			}
			continue

		case strings.HasPrefix(normalized, "up,"):
			nstr := strings.TrimPrefix(normalized, "up,")
			n, err := strconv.Atoi(nstr)
			if err == nil && n > 0 {
				start := len(result) - n
				if start < 0 {
					start = 0
				}
				for i := start; i < len(result); i++ {
					result[i] = upperWordPreserveFormat(result[i])
				}
				continue
			}
		}

		result = append(result, w)
	}
	return result
}
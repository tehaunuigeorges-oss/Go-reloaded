package fonction

import (
	"strconv"
	"strings"
)

func Cap(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		w := words[i]
		switch {
		case w == "(cap)":
			if len(result) > 0 {
				last := []rune(result[len(result)-1])
				if len(last) > 0 {
					result[len(result)-1] = strings.ToUpper(string(last[0])) + strings.ToLower(string(last[1:]))
				}
			}
			continue

		case w == "(cap," && i+1 < len(words):
			n, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err == nil && n > 0 && strings.HasSuffix(words[i+1], ")") {
				for j := max(len(result)-n, 0); j < len(result); j++ {
					runes := []rune(result[j])
					if len(runes) > 0 {
						result[j] = strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))
					}
				}
				i++
				continue
			}

		case strings.HasPrefix(w, "(cap,") && strings.HasSuffix(w, ")"):
			n, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(w, "(cap,"), ")"))
			if err == nil && n > 0 {
				for j := max(len(result)-n, 0); j < len(result); j++ {
					runes := []rune(result[j])
					if len(runes) > 0 {
						result[j] = strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))
					}
				}
				continue
			}
		}

		result = append(result, w)
	}
	return result
}

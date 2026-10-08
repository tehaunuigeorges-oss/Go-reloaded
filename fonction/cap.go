package fonction

import (
	"strconv"
	"strings"
)

func Cap(words []string) []string {
	var result []string

	for _, w := range words {
		switch {
		case w == "(cap)":
			if len(result) > 0 {
				last := []rune(result[len(result)-1])
				if len(last) > 0 {
					result[len(result)-1] = strings.ToUpper(string(last[0])) + string(last[1:])
				}
			}
			continue

		case strings.HasPrefix(w, "(cap,") && strings.HasSuffix(w, ")"):
			nstr := strings.TrimSuffix(strings.TrimPrefix(w, "(cap,"), ")")
			n, err := strconv.Atoi(nstr)
			if err == nil && n > 0 {
				start := max(len(result) - n, 0)
				for i := start; i < len(result); i++ {
					runes := []rune(result[i])
					if len(runes) > 0 {
						result[i] = strings.ToUpper(string(runes[0])) + string(runes[1:])
					}
				}
				continue
			}
		}

		result = append(result, w)
	}
	return result
}

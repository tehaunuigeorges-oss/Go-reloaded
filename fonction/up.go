package fonction

import (
	"strconv"
	"strings"
)

func Up(words []string) []string {
	var result []string

	for _, w := range words {
		switch {
		case w == "(up)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
			continue

		case strings.HasPrefix(w, "(up,") && strings.HasSuffix(w, ")"):
			nstr := strings.TrimSuffix(strings.TrimPrefix(w, "(up,"), ")")
			n, err := strconv.Atoi(nstr)
			if err == nil && n > 0 {
				start := len(result) - n
				if start < 0 {
					start = 0
				}
				for i := start; i < len(result); i++ {
					result[i] = strings.ToUpper(result[i])
				}
				continue
			}
		}

		result = append(result, w)
	}
	return result
}

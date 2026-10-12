package fonction

import (
	"strconv"
	"strings"
)

func Low(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		w := words[i]
		switch {
		case w == "(low)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
			continue

		case w == "(low," && i+1 < len(words):
			n, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err == nil && n > 0 && strings.HasSuffix(words[i+1], ")") {
				for j := max(len(result)-n, 0); j < len(result); j++ {
					result[j] = strings.ToLower(result[j])
				}
				i++
				continue
			}

		case strings.HasPrefix(w, "(low,") && strings.HasSuffix(w, ")"):
			n, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(w, "(low,"), ")"))
			if err == nil && n > 0 {
				for j := max(len(result)-n, 0); j < len(result); j++ {
					result[j] = strings.ToLower(result[j])
				}
				continue
			}
		}

		result = append(result, w)
	}
	return result
}

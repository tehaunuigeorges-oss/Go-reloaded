package fonction

import (
	"strings"
)

const Ponct = ".,!?:;"

func Ponctuation(words []string) []string {
	var result []string

	for _, w := range words {
		reste := strings.TrimLeft(w, Ponct)
		ponct := w[:len(w)-len(reste)]

		if ponct != "" && len(result) > 0 {
			result[len(result)-1] += ponct
			if reste != "" {
				result = append(result, reste)
			}
			continue
		}
		result = append(result, w)
	}
	return result
}

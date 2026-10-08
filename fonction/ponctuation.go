package fonction

import (
	"strings"
)

func Ponctuation(words []string) []string {
	var result []string

	for _, w := range words {
		if w == "," || w == "." || w == "!" || w == "?" || w == ":" || w == ";" || w == "..." || w == "!?" {
			if len(result) > 0 {
				result[len(result)-1] += w
				continue
			}
		}
		if strings.HasPrefix(w, ",") || strings.HasPrefix(w, ".") || strings.HasPrefix(w, "!") || strings.HasPrefix(w, "?") || strings.HasPrefix(w, ":") || strings.HasPrefix(w, ";") {
			if len(result) > 0 {
				runes := []rune(w)
				result[len(result)-1] += string(runes[0])
				rest := string(runes[1:])
				result = append(result, rest)
				continue
			}
		}
		result = append(result, w)
	}
	return result
}
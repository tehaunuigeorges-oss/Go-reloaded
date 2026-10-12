package fonction

import "strings"

func An(words []string) []string {
	result := make([]string, 0, len(words))

	for i, w := range words {
		if (w == "a" || w == "A") && i+1 < len(words) {
			next := strings.ToLower(words[i+1])
			if next != "" && strings.ContainsRune("aeiouh", []rune(next)[0]) {
				w += "n"
			}
		}
		result = append(result, w)
	}

	return result
}

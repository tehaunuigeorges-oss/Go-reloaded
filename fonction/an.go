package fonction

import ()


func An(words []string) []string {
	result := make([]string, 0, len(words))

	for _, w := range words {
		if w == "a" {
			runes := []rune(w)
			if len(runes) > 0 {
				switch runes[0] {
				case 'a', 'e', 'i', 'o', 'u', 'h':
					w = "an"
				}
			}
		}
		result = append(result, w)
	}

	return result
}
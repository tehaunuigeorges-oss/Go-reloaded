package fonction

import "strconv"

func Bin(words []string) []string {
	var result []string
	for _, w := range words {
		if w == "(bin)" && len(result) > 0 {
			lastWords := result[len(result)-1]
			if valeur, err := strconv.ParseInt(lastWords, 2, 64); err == nil {
				result[len(result)-1] = strconv.FormatInt(valeur, 10)
				continue
			}
		}
		result = append(result, w)
	}
	return result
}

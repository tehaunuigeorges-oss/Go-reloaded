package main

import "strconv"

func bin(words []string) []string {
	var result []string
	for _, w := range words {
		if w == "(bin)" {
			if len(result) > 0 {
				lastWords := result[len(result)-1]
				if valeur, err := strconv.ParseInt(lastWords, 2, 64); err == nil {
					result[len(result)-1] = strconv.FormatInt(valeur, 10)
				} else {
					result = append(result, w)
				}

			}
		}else{
			result = append(result, w)
		}
	}
	return result
}
package main

import (
	"strconv"
)

func hex(words []string) []string {
	var result []string
	for _, w := range words {
		if w == "(hex)" {
			if len(result) > 0 {
				lastwords := result[len(result)-1]
				if valeur, err := strconv.ParseInt(lastwords, 16, 64); err == nil {
					result[len(result)-1] = strconv.FormatInt(valeur, 10)
				} else {
					result = append(result, w)
				}
			} else {
				result = append(result, w)
			}
		} else {
			result = append(result, w)
		}
	}
	return result
}
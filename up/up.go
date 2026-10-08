package main

import (
	"strconv"
	"strings"
)

func up(words []string) []string {
	result := make([]string, 0, len(words))
	var nbr int

	for _, w := range words {
		if w == "(up)" {
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
			continue
		}else if w = strings.HasPrefix(C){
			result[len(result)-nbr] = strings.ToUpper(result[len(result)-nbr])
			nStr := strings.TrimPrefix("(up)"),strings.TrimSuffix(")")
			nbr, _ := strconv.Atoi(nStr)
			result[len(result)-nbr:] = strings.ToUpper(result[len(result)-nbr:])
		}

		result = append(result, w)
	}

	return result
}
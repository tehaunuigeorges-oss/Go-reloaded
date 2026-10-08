package fonction

import (
    "strconv"
    "strings"
)

func Low(words []string) []string {
    var result []string

    for _, w := range words {
        switch {
        case w == "(low)":
            if len(result) > 0 {
                result[len(result)-1] = strings.ToLower(result[len(result)-1])
            }
            continue

        case strings.HasPrefix(w, "(low,") && strings.HasSuffix(w, ")"):
            nstr := strings.TrimSuffix(strings.TrimPrefix(w, "(low,"), ")")
            n, err := strconv.Atoi(nstr)
            if err == nil && n > 0 {
                start := max(len(result) - n, 0)
                for i := start; i < len(result); i++ {
                    result[i] = strings.ToLower(result[i])
                }
                continue
            }
        }

        result = append(result, w)
    }
    return result
}
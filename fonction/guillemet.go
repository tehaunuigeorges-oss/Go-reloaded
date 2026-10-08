package fonction

func Guillemet(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		if words[i] != "'" && words[i] != `"` {
			result = append(result, words[i])
			continue
		}

		quote := words[i]
		close := i + 1
		for close < len(words) && words[close] != quote {
			close++
		}

		if close == len(words) {
			result = append(result, quote)
			continue
		}
		if close == i+1 {
			result = append(result, "''")
			i = close
			continue
		}

		result = append(result, "'"+words[i+1])
		for j := i + 2; j < close; j++ {
			result = append(result, words[j])
		}
		result[len(result)-1] += "'"
		i = close
	}

	return result
}

package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	words := strings.Fields(str)
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	keys := make([]string, 0, len(freq))

	for key := range freq {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	var result []string

	for i := 0; i < len(keys); i++ {
		if i < 10 {
			result = append(result, keys[i])
		}
	}

	return result
}

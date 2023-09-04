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

	sort.SliceStable(keys, func(i, j int) bool {
		if freq[keys[i]] == freq[keys[j]] {
			return strings.Compare(keys[i], keys[j]) <= 0
		}
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

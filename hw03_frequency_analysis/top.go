package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func Top10(input string) []string {
	if input == "" {
		return []string{}
	}
	words := strings.Fields(input)

	wordsCountMap := make(map[string]int)
	for _, word := range words {
		wordsCountMap[word]++
	}

	wordsResultCountSlice := make([]WordCount, 0, len(wordsCountMap))
	for word, count := range wordsCountMap {
		wordsResultCountSlice = append(wordsResultCountSlice, WordCount{
			Word:  word,
			Count: count,
		})
	}

	sort.SliceStable(wordsResultCountSlice, func(i, j int) bool {
		return wordsResultCountSlice[i].Word < wordsResultCountSlice[j].Word
	})
	sort.SliceStable(wordsResultCountSlice, func(i, j int) bool {
		return wordsResultCountSlice[i].Count > wordsResultCountSlice[j].Count
	})

	length := len(wordsResultCountSlice)
	if length > 10 {
		length = 10
	}

	result := make([]string, length)
	for i, word := range wordsResultCountSlice[:length] {
		result[i] = word.Word
	}

	return result
}

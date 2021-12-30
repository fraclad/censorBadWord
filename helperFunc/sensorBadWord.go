package helperFunc

import (
	"strings"
)

func CensorBadWord(mapToCheck map[string]struct{}, sentence string) string {
	words := strings.Split(sentence, " ")
	for idx, elem := range words {
		for val := range mapToCheck {
			if strings.TrimSpace(strings.ToLower(val)) == strings.TrimSpace(strings.ToLower(elem)) {
				words[idx] = strings.Repeat("*", len(elem))
			}
		}
	}
	finalSentence := strings.Join(words, " ")
	return (finalSentence)
}

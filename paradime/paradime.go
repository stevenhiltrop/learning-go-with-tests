package paradime

import "strings"

func Reverse(word string) string {
	chars := strings.Split(word, "")
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return strings.Join(chars, "")
}

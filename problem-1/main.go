package problem_1

import "strings"

func mergeAlternately(word1 string, word2 string) string {

	var builder strings.Builder

	for i := 0; i < max(len(word1), len(word2)); i++ {
		if i < len(word1) {
			builder.WriteByte(word1[i])
		}
		if i < len(word2) {
			builder.WriteByte(word2[i])
		}

	}

	return builder.String()

}

func main() {
	mergeAlternately("abc", "pqr")
}

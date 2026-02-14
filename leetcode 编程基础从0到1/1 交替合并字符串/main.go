package main

import (
	"fmt"
)

// 版本1
func mergeAlternately(word1 string, word2 string) string {
	var result []byte
	var i, j int
	for i < len(word1) && j < len(word2) {
		result = append(result, word1[i])
		i++
		result = append(result, word2[j])
		j++
	}
	if i < len(word1) {
		result = append(result, word1[i:]...)
	}
	if j < len(word2) {
		result = append(result, word2[j:]...)
	}
	return string(result)
}

func main() {
	fmt.Println(mergeAlternately("hello", "worlddddddd"))

}

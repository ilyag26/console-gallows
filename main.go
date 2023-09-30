package main

import (
	l "ilya/basic"
	"math/rand"
)

func main() {
	var word string
	words := []string{"golang", "java", "python", "ruby", "javascript"}
	min := 0
	max := len(words) - 1
	randWordNumber := rand.Intn(max-min) + min
	l.PrintRules()
	letterArray := l.TrimWord(words, randWordNumber)
	l.PrintWordByStep(words, letterArray, word, randWordNumber)
}

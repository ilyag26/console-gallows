package main

import (
	l "ilya/basic"
	"math/rand"
)

func main() {
	words := []string{"golang", "java", "python", "ruby", "javascript"}
	wordsExplanation := []string{"a compiled multi-threaded programming language developed internally by Google",
		"a strongly typed, object-oriented, general-purpose programming language developed by Sun Microsystems",
		"high-level general-purpose programming language with dynamic strong typing and automatic memory management",
		"dynamic, reflective, interpreted high-level programming language",
		"multi-paradigm programming language. Supports object-oriented, imperative and functional styles"}
	min := 0
	max := len(words) - 1
	randWordNumber := rand.Intn(max-min) + min
	l.PrintRules()
	letterArray := l.TrimWord(words, randWordNumber)
	l.PrintWordByStep(words, wordsExplanation, letterArray, randWordNumber)
}

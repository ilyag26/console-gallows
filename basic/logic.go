package basic

import (
	"fmt"
	"strings"
)

func PrintRules() {
	fmt.Println("-----GO GALLOWS----",
		"\nGame rules: you need to guess the hidden word!",
		"\nWrite word that you think hidden and try to win!",
		"\n-----------")
}

func TrimWord(words []string, randNumber int) []string {
	return strings.Split(words[randNumber], "")
}

func PrintWordByStep(hiddenWord []string, letterArray []string, word string, numberRand int) {
	var wordClue string
	for i := 0; i != len(letterArray); i++ {
		wordClue += letterArray[i]
		fmt.Println("Try to guess te word -> ", "Following words: '",
			wordClue, "' |", "Number of letters: ", len(letterArray))
		fmt.Scan(&word)
		if word == hiddenWord[numberRand] {
			fmt.Println("DONE! -> YOU WIN!")
			break
		} else if wordClue == hiddenWord[numberRand] {
			fmt.Println("You lose!")
			break
		} else {
			fmt.Println("NOPE! Try again")
		}
	}
}

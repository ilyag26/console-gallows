package basic

import (
	"fmt"
	"slices"
	"strings"
)

var (
	letter            string
	idLetter          int
	checkGame         bool
	iterationGameLose int
	iterationGameWin  int
	attemps           int
)

func PrintRules() {
	fmt.Println("-----GO GALLOWS----",
		"\nGame rules: you need to guess the hidden word!",
		"\nWrite word that you think hidden and try to win!",
		"\n-----------")
}
func gallowsPrint(attemps int) {
	switch {
	case attemps == 5:
		fmt.Print("|--------\n|       |\n|       O\n|\n|\n|\n")
	case attemps == 4:
		fmt.Print("|--------\n|       |\n|       O\n|       |\n|\n|\n")
	case attemps == 3:
		fmt.Print("|--------\n|       |\n|       O\n|      /|\\\n|\n|\n")
	case attemps == 2:
		fmt.Print("|--------\n|       |\n|       O\n|      /|\\\n|	|\n|\n")
	case attemps == 1:
		fmt.Print("|--------\n|       |\n|       O\n|      /|\\\n|	|\n|      / \\\n")
	}
}
func maskingWord(lengthWord int) []string {
	var mask []string
	for i := 0; i != lengthWord; i++ {
		mask = append(mask, "*")
	}
	return mask
}

func TrimWord(words []string, randNumber int) []string {
	return strings.Split(words[randNumber], "")
}

func PrintWordByStep(hiddenWord []string, wordsExplanation []string, letterArray []string, numberRand int) {
	maskedWord := maskingWord(len(letterArray))
	attemps = 6
	fmt.Println("What is this? -> ", wordsExplanation[numberRand])
	for !checkGame {
		fmt.Println("Try to guess the word -> ", "Number of letters in this word: ", len(letterArray), "Remainings Attemps: ", attemps)
		fmt.Print("Word is: ")
		for _, value := range maskedWord {
			fmt.Printf("%s", value)
		}
		fmt.Print("\nEnter letter: ")
		fmt.Scan(&letter)
		if slices.Contains(letterArray, letter) {
			if iterationGameWin != len(letterArray)-1 {
				idLetter = strings.Index(hiddenWord[numberRand], letter)
				maskedWord[idLetter] = letter
				iterationGameWin++
				fmt.Printf("good, there is '%s'\n", letter)
			} else {
				fmt.Printf("You win. Word is '%s'\n", hiddenWord[numberRand])
				break
			}
		} else {
			if iterationGameLose != 5 {
				fmt.Printf("NO!, there is not '%s'\n", letter)
				iterationGameLose++
				attemps--
				gallowsPrint(attemps)
			} else {
				fmt.Printf("You lose. Word was '%s'\n", hiddenWord[numberRand])
				break
			}
		}
	}
}

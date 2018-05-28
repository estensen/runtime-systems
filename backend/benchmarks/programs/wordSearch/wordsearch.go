package wordsearch

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var wordMap = make(map[string]int)
var wordList = []string{}

func WordSearchWithMap(word string) {
	lowerStringWord := strings.ToLower(word)
	readTextFile(true)
	if _, ok := wordMap[lowerStringWord]; ok {
		fmt.Printf("Word: %s exists in story\n", word)
	} else {
		fmt.Printf("Word: %s does not exist in story\n", word)
	}
}

func WordSearchWithList(word string) {
	lowerStringWord := strings.ToLower(word)
	readTextFile(false)
	wordExists := false
	for _, textword := range wordList {
		if textword == lowerStringWord {
			fmt.Printf("Word: %s exists in story\n", word)
			wordExists = true
			break
		}
	}
	if wordExists {
		fmt.Printf("Word: %s does not exist in story\n", word)
	}

}

func stringToAllAlpha(text string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal("Unable to convert string to only alpha")
	}
	return reg.ReplaceAllString(text, "")
}

func readTextFile(useMap bool) {
	text, err := os.Open("programs/wordSearch/story.txt")
	if err != nil {
		panic("Unable to read textfile")
	}
	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)
	if useMap {
		for scanner.Scan() {
			lowerStringWord := stringToAllAlpha(strings.ToLower(scanner.Text()))
			wordMap[lowerStringWord] = wordMap[lowerStringWord] + 1
		}
	} else {
		for scanner.Scan() {
			lowerStringWord := stringToAllAlpha(strings.ToLower(scanner.Text()))
			fmt.Println(lowerStringWord)
			wordList = append(wordList, lowerStringWord)

		}
	}
}

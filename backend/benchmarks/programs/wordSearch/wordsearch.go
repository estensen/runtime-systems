package wordSearch

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
	readTextFile(true)
	if checkifWordExistsMap(word) {
		fmt.Printf("Word: %s exists in story\n", word)
	} else {
		fmt.Printf("Word: %s does not exist in story\n", word)
	}
}

func checkifWordExistsMap(word string) bool {
	lowerStringWord := strings.ToLower(word)
	if _, ok := wordMap[lowerStringWord]; ok {
		return true
	}
	return false
}

func WordSearchWithList(word string) {
	readTextFile(false)
	if checkIfWordExistsList(word) {
		fmt.Printf("Word: %s exists in story\n", word)
	} else {
		fmt.Printf("Word: %s does not exist in story\n", word)
	}
}

func checkIfWordExistsList(word string) bool {
	lowerStringWord := strings.ToLower(word)
	readTextFile(false)
	for _, textword := range wordList {
		if textword == lowerStringWord {
			return true
		}
	}
	return false
}

func stringToAllAlpha(text string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal("Unable to convert string to only alpha")
	}
	return reg.ReplaceAllString(text, "")
}

func readTextFile(useMap bool) {
	text, err := os.Open("./benchmarks/programs/wordSearch/story.txt")
	if err != nil {
		panic("Unable to read story")
	}
	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)
	if useMap {
		for scanner.Scan() {
			lowerStringWord := stringToAllAlpha(strings.ToLower(scanner.Text()))
			wordMap[lowerStringWord]++
		}
	} else {
		for scanner.Scan() {
			lowerStringWord := stringToAllAlpha(strings.ToLower(scanner.Text()))
			wordList = append(wordList, lowerStringWord)

		}
	}
}

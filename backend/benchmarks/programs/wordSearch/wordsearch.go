package wordsearch

import (
	"fmt"
	"io/ioutil"
)

var wordMap = make(map[string]int)
var wordList = make([]string, 0)

func WordSearchWithMap(word string) {
	readTextFile(true)
	if _, ok := wordMap[word]; ok {
		fmt.Println("Word: %s exists in story", word)
	}
	fmt.Println("Word: %s does not exist in story", word)
}

func WordSearchWithList(word string) {
	readTextFile(false)
	for _, textword := range wordList {
		if textword == word {
			fmt.Println("Word: %s exists in story", word)
		}
	}
	fmt.Println("Word: %s does not exist in story", word)
}

func readTextFile(useMap bool) {
	text, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		panic("Unable to read textfile")
	}
	if useMap {
		for _, word := range text {
			wordMap[string(word)] = wordMap[string(word)] + 1
		}
	} else {
		for i, word := range text {
			wordList[i] = string(word)
		}
	}
}

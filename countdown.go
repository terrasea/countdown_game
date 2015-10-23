package main

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"strings"
	"os"
)

func loadDict() ([]string, error) {
	body, err := ioutil.ReadFile("dict.txt")
	if err != nil {
		return nil, err
	}

	return strings.Split(string(body), "\n"), nil
}

func checkWord(letters []byte, word []byte) (int) {
	for i := 0; i < len(word); i++ {
		pos := bytes.IndexByte(letters, word[i])

		if pos < 0 {
			return pos
		} else {
			letters[pos] = letters[len(letters)-1]
			letters = letters[:len(letters)-1]
		}
	}

	return len(word)
}


func getBestWord(words []string, letters string) (int, []string) {
	var maxlength = 0
	var currentwords []string
	for _, word := range words {
		var length = checkWord([]byte(letters), []byte(word))
		if length >= maxlength {
			if length > maxlength {
				currentwords = make([]string, 0)
			}
			maxlength = length
			currentwords = append(currentwords, word)
		}
	}

	return maxlength, currentwords
}

func main() {
	argsWithoutProg := os.Args[1:]

	var words, _ = loadDict()
	var letters = argsWithoutProg[0]
	var maxlength, currentword = getBestWord(words, letters)

	fmt.Print(letters, ", ", maxlength, ", ", currentword, "\n")
}

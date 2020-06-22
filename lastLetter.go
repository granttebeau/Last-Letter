package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Need a text file")
		os.Exit(1)
	}

	content, err := readLines(args[0])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(sequence(content))
}

func sequence(words []string) []string {
	return sequenceHelper(words, 0)
}

func sequenceHelper(words []string, index int) []string {
	if index == len(words) {
		return []string{}
	}
	word := words[index]

	longestSequence := getLongestSequence(word, words)

	index++

	return longestSlice(longestSequence, sequenceHelper(words, index))
}

func getLongestSequence(word string, words []string) []string {
	var wordsCopy = make([]string, len(words))
	copy(wordsCopy, words)

	var sequence []string
	sequence = append(sequence, word)
	lastLetter := string([]rune(word)[len(word)-1])
	for {
		nextWord := findNextWord(word, lastLetter, wordsCopy)
		if nextWord == "NO MORE WORDS" {
			break
		}
		if nextWord != word {
			sequence = append(sequence, nextWord)
			lastLetter = string([]rune(nextWord)[len(nextWord)-1])
			indexOfWord := getIndex(wordsCopy, nextWord)
			wordsCopy = append(wordsCopy[:indexOfWord], wordsCopy[indexOfWord+1:]...)
		}
	}
	return sequence
}

func findNextWord(origWord string, letter string, words []string) string {
	for _, word := range words {
		if string([]rune(word)[0]) == letter && word != origWord {
			return word
		}
	}
	return "NO MORE WORDS"
}

func longestSlice(s1 []string, s2 []string) []string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getIndex(words []string, w string) int {
	for index, word := range words {
		if w == word {
			return index
		}
	}
	return 0
}

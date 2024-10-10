package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
    result := TopWords("declaration_of_independence.txt", 10, 5)
    for _, wc := range result {
        fmt.Println(wc)
    }
}

func TopWords(path string, numWords int, charThreshold int) []WordCount {
	file, err := os.Open(path)
	checkError(err)
	defer file.Close()
	wordCounts := make(map[string]int)
	var wordCountSlice []WordCount

	for {
		line, err := readLine(file)
		if err != nil {
			break
		}

		words := strings.Fields(strings.ToLower(line))

		for i, word := range words {
			words[i] = removePunctuations(word)
		}

		for _, word := range words {
			if len(word) >= charThreshold {
				wordCounts[word]++
			}
		}
	}

	for word, count := range wordCounts {
		wordCountSlice = append(wordCountSlice, WordCount{word, count})
	}

	sortWordCounts(wordCountSlice)


	return wordCountSlice[:numWords]
}

// Helper function to read a line from a file
func readLine(file *os.File) (string, error) {
	var line string
	var err error
	var buf [1]byte

	for {
		_, err = file.Read(buf[:])
		if err != nil {
			return "", err
		}
		if buf[0] == '\n' {
			break
		}
		line += string(buf[0])
	}

	return line, nil
}

func removePunctuations(str string) string {
	var newLine string
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			newLine += string(char)
		}
	}
	return newLine
}

type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}


func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

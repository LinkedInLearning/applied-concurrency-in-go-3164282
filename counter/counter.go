package counter

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type Statistics struct {
	counts map[string]int
}

func New() *Statistics {
	var c Statistics
	c.counts = make(map[string]int)
	return &c
}

// AddWord a new word to the counter or
// Increments word count if word already exists in the counter
func (s *Statistics) AddWord(w string) {
	// strip common punctuation from the word
	word := strings.Trim(w, ",.!?-\"")
	// make sure we have not ended up with an empty word
	if _, ok := CommonWords[strings.ToLower(word)]; !ok && word != "" {
		// add the word to the map
		s.counts[word] = s.counts[word] + 1
	}
}

// wordPrint is a nested struct contains
// both a word and its count for sorting and printing
type wordPrint struct {
	word  string
	count int
}

// Prints the count stats that are larger than 1 to the console
func (s Statistics) PrintStats() {
	// convert the map to list so we can order it
	var wordList []wordPrint
	for word, count := range s.counts {
		wordList = append(wordList, wordPrint{word, count})
	}
	// sort the word list with higher count first
	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].count > wordList[j].count
	})
	// print the stats for all the words in a table
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer writer.Flush()
	for _, w := range wordList {
		// print only words that occur more than twice to reduce noise
		if w.count > 2 {
			fmt.Fprintf(writer, "%s\t%d\n", w.word, w.count)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
		text := insertAString()
		if searchInStringOccurrences(text) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}

}

func searchInStringOccurrences(text string) bool {
	st := strings.TrimSpace(strings.ToLower(text))
	foundI := strings.Index(st, "i") == 0
	foundN := strings.LastIndexAny(st, "n") == len(st)-1
	foundA := strings.Index(st, "a") > 0 && strings.IndexAny(st, "a") < len(st)-1
	return foundI && foundN && foundA
}

func insertAString() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a string:")
	text, err := r.ReadString('\n')
	if err != nil {
		panic("Something wrong happened!")
	}
	return text
}

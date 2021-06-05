package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sl := make([]int, 0,3)
	exit := "0"
	for !strings.EqualFold(exit, "X") {
		st := insertANumberFromCommandLine()
		if len(st) > 0 { //not empty string check
			exit = st
			num, err := strconv.Atoi(st)
			if err != nil && !strings.EqualFold(st,"X"){
				fmt.Println("something wrong with what you entered, it's not a number...Try again or please enter 'X' to exit the program!")
			} else {
				sl = append(sl, num)
				sort.Ints(sl)
				fmt.Println("Sorted array =>", sl)
			}
		} else {
			fmt.Println("Empty character cannot be entered...Try again or please enter 'X' to exit the program!!")
		}
	}

}

func insertANumberFromCommandLine() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Please insert an int number or Press 'X' to exit the program:")
	st, err := r.ReadString('\n')
	if err != nil {
		panic("Something wrong happened!")
	}
	return strings.TrimSpace(st)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	st := insertANumberFromCommandLine()
	truncateAndPringAFloat(st)
	st2 := insertANumberFromCommandLine()
	truncateAndPringAFloat(st2)
}

func insertANumberFromCommandLine() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Please insert a float number:")
	st, err := r.ReadString('\n')
	if err != nil {
		panic("Something wrong happened!")
	}
	return st
}

func truncateAndPringAFloat(st string) {
	st = strings.Replace(st, "\n", "", -1)
	if fl, err := strconv.ParseFloat(st, 64); err == nil {
		fmt.Printf("My float is -> %f and is Truncated Integer is ->  %d  \n", fl, int(fl))
		fmt.Printf("My float in decimal precision 2 is -> %.2f and is round version Integer is ->  %v\n", fl, math.Round(fl))
	} else {
		fmt.Printf("Not a Float - Please Re-Enter \n")
	}
}

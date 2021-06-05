package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.*/

func main() {

	numbers := make([]int, 0)
	numbers = insertASliceOfIntegers(numbers)
	fmt.Printf("Original slice of integers-> %v with size %d and capacity %d \n", numbers, len(numbers), cap(numbers))
	ln := len(numbers)
	if ln > 1 {
		bubbleSort(numbers, ln)
	}
	fmt.Printf("Sorted Slice of integers-> %v\n", numbers)
}

func bubbleSort(numbers []int, ln int) {
	for i := 0; i < ln; i++ {
		for j := 0; j < ln-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				SwapNum(numbers, j)
			}
		}
	}
}

func SwapNum(numbers []int, j int) {
	temp := numbers[j]
	numbers[j] = numbers[j+1]
	numbers[j+1] = temp
}

func insertASliceOfIntegers(numbers []int) []int {
	readCmdLine := bufio.NewReader(os.Stdin)
	fmt.Println("Remember max 10 elements!!")

	for len(numbers) < 10 {
		fmt.Println("Insert a new integer or Press 'X' to exit ==>")
		num, err := readCmdLine.ReadString('\n')
		checkError(err)
		num = strings.Trim(num, "\n")
		num = strings.TrimSpace(num)
		if strings.ToUpper(strings.TrimSpace(num)) == "X" {
			break
		} else {
			n, err2 := strconv.Atoi(num)
			checkError(err2)
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Something wrong happened, Error is: {}", err)
	}
}

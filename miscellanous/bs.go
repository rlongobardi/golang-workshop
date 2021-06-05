package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = tmp
}

func BubbleSort(numbers []int) {
	Swapped := true
	for Swapped {
		Swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				Swapped = true
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter X to exit")
	fmt.Printf("Please enter a list of integers (space-separated): ")

	for (scanner.Scan()) {
		var strVal string = scanner.Text()
		if strVal == "X" {
			os.Exit(0)
		}

		var numlist = strings.Split(strVal, " ")
		var intSlice = make([]int, 0)
		for _, i := range numlist {
			j, _ := strconv.Atoi(i)
			intSlice = append(intSlice, j)
		}
		if (len(intSlice) > 10) {
			fmt.Println("Please enter a maximum of 10 integers")
			fmt.Printf("Please enter a list of integers (space-separated): ")
			continue
		}
		BubbleSort(intSlice)
		fmt.Println(intSlice)
		fmt.Printf("Please enter a list of integers (space-separated): ")
	}
}
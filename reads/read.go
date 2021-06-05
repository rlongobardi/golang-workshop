package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.*/

const SingleSpace = " "

type Person struct {
	fname string
	lname string
}

func main() {
	readerCmdLn := bufio.NewReader(os.Stdin)
	people := make([]Person, 0, 3)
	//get filename from command line
	fmt.Println("Please enter the filename of the text file:")
	filename, err := readerCmdLn.ReadString('\n')
	check(err)
	filename = strings.Trim(filename, "\n")
	//file handling
	file, errFile := os.Open(filename)
	check(errFile)
	defer file.Close()
	// read file
	people = parseFileLineByLine(file, people)
	printListOfPeople(people)

}

func parseFileLineByLine(file *os.File, people []Person) []Person {
	scanFile := bufio.NewScanner(file)
	for scanFile.Scan() {
		line := strings.Split(scanFile.Text(), SingleSpace)
		var person Person
		person.fname = line[0]
		person.lname = line[1]
		people = append(people, person)
	}
	return people
}

func printListOfPeople(people []Person) {
	for _, p := range people {
		fmt.Printf("Firstname is %v and Lastname is %v \n", p.fname, p.lname)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

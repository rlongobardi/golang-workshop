package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}


func main() {

	people := make([]Name,0)
	var filename string
	fmt.Println("Enter the filename:")

	_, err1 := fmt.Scanln(&filename)
	if err1 != nil && err1 != io.EOF {
		log.Fatal("Error: ", err1)
	}
	filename = strings.Trim(filename,"\n")

	f, ferr :=os.Open(filename)
	checkError(ferr)
	scanLine := bufio.NewScanner(f)
	for scanLine.Scan(){
		line := scanLine.Text()
		var person Name
		splitLine := strings.Split(line, " ")
		person.fname, person.lname = splitLine[0], splitLine[1]
		people = append(people, person)
	}

	for i, p := range people{
		fmt.Printf("%d %s %s \n",i+1,p.fname, p.lname)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Something bad happened, please try again, error -> ",err)
	}
}
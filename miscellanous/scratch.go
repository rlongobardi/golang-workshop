package main

import "fmt"
import "os"
import "sort"
import "strconv"

func main() {
	slice := make([]int, 3)
	insertedNumber := ""
	index := 0
	for true {
		fmt.Printf("Please, insert a integer number ---> ")
		fmt.Scan(&insertedNumber)
		if(insertedNumber == "X"){
			fmt.Scan("Inserted character X, exit the loop!")
			break
		}else{
			if(index>2){
				slice = append(slice,1)
			}
			insertedNumberInt, err := strconv.Atoi(insertedNumber)
			if(err != nil){
				fmt.Println(err)
				os.Exit(1)
			}
			slice[index] = insertedNumberInt
			sort.Ints(slice)
			fmt.Println(slice)
			index += 1
		}
	}
}

package main

import (
	"fmt"
)
//how to create IOTA
const (
	Monday int = iota +1
	Tuesday
	Wednesday
	)

var (
	i = 23
	alex = "Nemo"
)
func main() {
/*	var raw_string string

	fmt.Scanln(&raw_string)

	a_is_contained := strings.Contains(raw_string, "a")
	starts_with_i := raw_string[0] == 'i'
	finishes_with_n := raw_string[len(raw_string)-1] == 'n'
	if a_is_contained && starts_with_i && finishes_with_n {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
*/
	fmt.Printf("constants %v and %v and %v\n",Monday,Tuesday, Wednesday)
	fmt.Printf("variables %d and %s\n",i,alex)

}

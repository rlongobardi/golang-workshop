package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)

	flt, err := strconv.ParseFloat(str, 64)
	fmt.Println("truncated", int(flt))

	if err != nil {
		fmt.Println(err)
	}
}
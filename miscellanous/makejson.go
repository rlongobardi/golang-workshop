package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	mapJson := make(map[string]string)
	getNameAndAddress(mapJson)

	fmt.Println(" mapjson->", mapJson)
	jsonByte, err := json.Marshal(mapJson)
	if err != nil {
		panic("Something wrong happened when trying to marshalling!")
	}
	jsonPrint := string(jsonByte)
	fmt.Println("Json is ->", jsonPrint)

}

func getNameAndAddress(myMap map[string]string) {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a name: ")
	name, err := r.ReadString('\n')
	if err != nil {
		panic("Something wrong happened!")
	}
	fmt.Println("Please enter the address: ")
	address, err2 := r.ReadString('\n')
	if err2 != nil {
		panic("Something wrong happened!")
	}
	myMap["name"] = strings.Trim(name,"\n")
	myMap["address"] = strings.Trim(address,"\n")
}

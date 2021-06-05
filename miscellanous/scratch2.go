/*package main

import (
	"encoding/json"
	"fmt"
)
*/
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	for {
//		fmt.Println("Enter a float value [Ctrl-C to Exit Program]:")
//		text, _ := reader.ReadString('\n')
//		text = strings.Replace(text, "\n", "", -1)
//		if fl, err := strconv.ParseFloat(text, 64); err == nil {
//			fmt.Printf("Float = %f :: Truncated Integer =  %d  \n", fl, int(fl))
//		} else {
//			fmt.Printf("Not a Float - Please Re-Enter \n")
//		}
//
//	}
//}

package main

import (
	"encoding/json"
	"fmt"
)


func main(){
	nameAndAddressMap := make(map[string]string)
	fmt.Println("Enter name")
	var name string
	fmt.Scan(&name)
	fmt.Println("Enter address")
	var address string
	fmt.Scan(&address)
	nameAndAddressMap["name"] = name
	nameAndAddressMap["address"] = address

	jsonObject, _ := json.Marshal(nameAndAddressMap)
	fmt.Println("Created Json Object", string(jsonObject))
}

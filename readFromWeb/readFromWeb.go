package readFromWeb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func ReadTextUrlFromWeb() {
	var url = "http://services.explorecalifornia.org/json/tours.php"
	resp, err := http.Get(url)
	CheckError(err)

	defer resp.Body.Close()
	fmt.Printf("The response type: %T\n", resp)
	bytes, err2 := ioutil.ReadAll(resp.Body)
	CheckError(err2)

	content := string(bytes)
	tours := ToursFromJson(content)
	//fmt.Println(tours)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, 0)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ToursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	CheckError(err)
	var tour Tour
	for decoder.More() {
		decoder.Decode(&tour)
		tours = append(tours, tour)
	}
	return tours
}

package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	PageNumber int `json:"page"`
	Fruits []string `json:"fruits"`
}

func encodeUsage() {
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapA)
	fmt.Println("Json encode usage is as below:")
	fmt.Println(string(mapB))
}

func decodeUsage() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response{}

	//While decoding the json byte using unmarshal.
	//the first argument is the json byte
	//the second argument is the address of the response type struct where we want the json to be mapped to
	json.Unmarshal([]byte(str), &res)

	fmt.Println("Json decode usage is as below:")
	fmt.Println(res)	//res type is defined in the struct at the top.
	fmt.Printf("Page is %v, fruits are %v", res.PageNumber, res.Fruits)
}

func main() {
	encodeUsage()
	decodeUsage()
}

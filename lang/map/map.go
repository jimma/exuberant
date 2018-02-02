package main

import "fmt"

func main() {
	var stringMap map[string]string
	stringMap = make(map[string]string)
	stringMap["key1"] = "value1"
	fmt.Printf("stringMap %s \n", stringMap)
	_, ok := stringMap["key2"]
	fmt.Printf("key2 exsits ? %v \n", ok)

}

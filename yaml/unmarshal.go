/*
Method that used to convert json(Byte data) into Struct
String in Json/YAML format ---> byte data
Byte data --> UnMarshal --> GO Object(struct)
*/

package main

import (
	"encoding/json"
	"fmt"
)

/*
We will create Response Struct that we will use to match byte code
after Unmarshal to display data.
*/
type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Example JSON string
	empJsonData := `{"Name":"George Smith","Age":30,"Address":"Newyork, USA"}`
	// JSON string into byte code
	empBytes := []byte(empJsonData)
	var emp Response
	json.Unmarshal(empBytes, &emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Address)
}

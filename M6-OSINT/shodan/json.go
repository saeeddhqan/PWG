
package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name 	string
	Family 	string
	Languages 	[]string
}


func main() {
	group := Person{Name: "John", Family: "Doe", Languages: []string{"C", "Java", ".."}}
	marshal, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
	// Unmarshal
	var json2struct Person
	// json.Unmarshal(marshal, &json2struct)
	json.Unmarshal([]byte(`{"Name": "Darlin", "Family": "Nilrad", "Languages": ["C", "C++", ".."]}`), &json2struct)
	fmt.Println(json2struct.Name, json2struct.Family)
}
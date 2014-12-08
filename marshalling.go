package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	p1 := Person{Name: "Homer"}
	jsonP1, _ := json.Marshal(p1)
	fmt.Printf("p1: %s\n", jsonP1)

	var p2 Person
	jsonP2 := "{\"name\":\"Bart\"}"
	json.Unmarshal([]byte(jsonP2), &p2)
	fmt.Printf("p2: %+v\n", p2)
}

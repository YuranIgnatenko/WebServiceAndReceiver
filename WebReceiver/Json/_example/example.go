package main

import (
	"Json"
	"fmt"
)

type Data struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	Flag bool   `json:"flag"`
}

/*
	for run example :
	cd /your_path/Json/_example
	go run example.go
*/
func main() {
	filename := "data.json"
	data := Data{}
	err := Json.JsonToStruct(filename, &data)
	fmt.Println(err)
	fmt.Printf("%#v\n", data)
}

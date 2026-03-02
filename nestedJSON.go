package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Customers string `json:"customers"`
	Items     string `json:"items`
}

func main() {
	o1 := Order{Customers: "Bintang", Items: "Buku"}
	jsonData, _ := json.Marshal(o1)
	fmt.Println("Hasil JSON:", string(jsonData))
}

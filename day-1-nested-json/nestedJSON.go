package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	SecretCode string `json:"-"`
}

type Items struct {
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	ID       int      `json:"id"`
	Customer Customer `json:"customer"`
	Items    []Items  `json:"items"`
}

func main() {
	//Marshal
	p1 := Order{
		ID: 101,
		Customer: Customer{
			Name:       "Bintang",
			Email:      "bintang@gmail.com",
			SecretCode: "AKUCANTIK",
		},
		Items: []Items{
			{ProductName: "Novel", Price: 200000, Quantity: 1},
			{ProductName: "Pencil", Price: 2000, Quantity: 2},
		},
	}

	jsonData, _ := json.Marshal(p1)
	fmt.Println("Hasil JSON:", string(jsonData))

	//Unmarshal
	jsonInput := `{
		"id":102,
		"customer":{
			"name":"Kirana",
			"email":"kirana@gmail.com"
		},
		"items":[
			{"product_name":"Melon Tea", "price":10000,"quantity":1},
			{"product_name":"Lemon Tea", "price":10000,"quantity":1}
		]
	}`
	var p2 Order
	err := json.Unmarshal([]byte(jsonInput), &p2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Hasil Struct:%+v\n", p2)
}

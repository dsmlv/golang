package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func MarshalProduct(p Product) ([]byte, error) {
	return json.Marshal(p)
}

func UnmarshalProduct(data []byte) (Product, error) {
	var p Product
	err := json.Unmarshal(data, &p)
	return p, err
}

func main() {
	product := Product{
		Name:     "Laptop",
		Price:    999.99,
		Quantity: 5,
	}

	jsonData, err := MarshalProduct(product)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	fmt.Println("JSON:", string(jsonData))

	newProduct, err := UnmarshalProduct(jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	fmt.Println("Decoded product:", newProduct)
}

package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	// This fake api was created with typicode json-server. For more detailed information;
	// https://github.com/typicode/json-server

	//project.GetAllProducts()

	project.AddProduct()
	products, _ := project.GetAllProducts()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
}
